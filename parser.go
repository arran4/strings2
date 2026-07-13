package strings2

import (
	"strings"
	"unicode"
)

// Parse parses the input string into a slice of Words based on detection or provided options.
// It follows the pipeline: String -> SubParts -> Parts -> Words.
//
// opts can be:
// - ParserOption interface
// - Partitioner function
// - PartitionerConfig
// - ParserSmartAcronyms bool
func Parse(input string, opts ...any) ([]Word, error) {
	// Level 5: Scan
	subs, stats := StringToSubParts(input)

	p := &ParserConfig{
		SmartAcronyms: true,
		NumberMode:    NumberModeNone,
	}

	var subPartMappers []SubPartMapper
	var partMappers []PartMapper
	var wordMappers []WordMapper

	for _, opt := range opts {
		switch o := opt.(type) {
		case Partitioner:
			p.Partitioner = o
		case PartitionerConfig:
			p.Partitioner = NewPartitioner(o)
		case ParserOption:
			o.Apply(p)
		case SubPartMapper:
			subPartMappers = append(subPartMappers, o)
		case PartMapper:
			partMappers = append(partMappers, o)
		case WordMapper:
			wordMappers = append(wordMappers, o)
		}
	}

	for _, m := range subPartMappers {
		if m != nil {
			subs = m(subs)
		}
	}

	// Level 4: Partition
	// If partitioner is not set, try to detect
	partitioner := p.Partitioner
	if partitioner == nil {
		partitioner = DetectPartitioner(stats, p)
	}

	parts := SubPartsToParts(subs, partitioner)

	for _, m := range partMappers {
		if m != nil {
			parts = m(parts)
		}
	}

	// Level 3: Words
	words := PartsToWords(parts, p)

	for _, m := range wordMappers {
		if m != nil {
			words = m(words)
		}
	}

	return words, nil
}

// ParserConfig holds configuration for the parsing pipeline.
type ParserConfig struct {
	Partitioner Partitioner
	// SmartAcronyms controls whether all-uppercase words (longer than 1 char)
	// should be treated as AcronymWord instead of UpperCaseWord.
	// Defaults to true.
	SmartAcronyms bool
	// NumberMode controls how numbers are handled during word splitting.
	NumberMode NumberMode
	// EmitEmpty controls whether empty parts are emitted for delimiters
	EmitEmpty bool
}

// NumberMode defines the strategy for handling numbers during parsing.
type NumberMode int

const (
	// NumberModeNone does not perform any special number splitting.
	NumberModeNone NumberMode = iota
	// NumberModeSplitAlways splits on any transition between a letter and a digit.
	NumberModeSplitAlways
	// NumberModeMergeWithWord treats digits as compatible with both preceding and succeeding lowercase letters,
	// preventing splits like 123test -> 123-test.
	NumberModeMergeWithWord
	// NumberModeTreatAsLowercase treats digits exactly as if they were lowercase letters for boundary detection.
	NumberModeTreatAsLowercase
)

// ParserOption configures the parser.
type ParserOption interface {
	Apply(*ParserConfig)
}

type funcParserOption func(*ParserConfig)

func (f funcParserOption) Apply(p *ParserConfig) { f(p) }

// ParserSmartAcronyms is a typed option for SmartAcronyms configuration.
// It allows passing a boolean-like type directly to Parse.
type ParserSmartAcronyms bool

func (b ParserSmartAcronyms) Apply(p *ParserConfig) {
	p.SmartAcronyms = bool(b)
}

// ParserEmitEmpty is a typed option for EmitEmpty configuration.
type ParserEmitEmpty bool

func (b ParserEmitEmpty) Apply(p *ParserConfig) {
	p.EmitEmpty = bool(b)
}

// WithPartitioner sets a specific partitioner strategy.
func WithPartitioner(pt Partitioner) ParserOption {
	return funcParserOption(func(p *ParserConfig) {
		p.Partitioner = pt
	})
}

// WithSmartAcronyms enables or disables smart acronym detection.
func WithSmartAcronyms(enabled bool) ParserOption {
	return funcParserOption(func(p *ParserConfig) {
		p.SmartAcronyms = enabled
	})
}

// WithNumberSplitting enables or disables splitting on letter-digit boundaries.
// It is equivalent to WithNumberMode(NumberModeSplitAlways) when true, and WithNumberMode(NumberModeNone) when false.
func WithNumberSplitting(enabled bool) ParserOption {
	return funcParserOption(func(p *ParserConfig) {
		if enabled {
			p.NumberMode = NumberModeSplitAlways
		} else {
			p.NumberMode = NumberModeNone
		}
	})
}

// WithNumberMode sets the specific number splitting mode.
func WithNumberMode(mode NumberMode) ParserOption {
	return funcParserOption(func(p *ParserConfig) {
		p.NumberMode = mode
	})
}

// DetectPartitioner uses stats to guess the best partitioner.
// config is optional, if provided it uses settings like NumberSplitting.
func DetectPartitioner(stats Stats, config ...*ParserConfig) Partitioner {
	delimiters := make(map[rune]bool)
	if stats.Spaces > 0 {
		delimiters[' '] = true
	}

	// Add known delimiters if present
	known := []rune{'_', '-', '+', '/', '\\', '|', ',', ';', ':'}
	for _, r := range known {
		if stats.SymbolCounts[r] > 0 {
			delimiters[r] = true
		}
	}

	// Check for dots.
	// If we have spaces, dots might be punctuation (end of sentence).
	// If no spaces, dots are likely delimiters (user.id).
	if stats.SymbolCounts['.'] > 0 {
		if stats.Spaces == 0 {
			// likely delimiter
			delimiters['.'] = true
		}
	}

	numberMode := NumberModeNone
	emitEmpty := false
	if len(config) > 0 && config[0] != nil {
		numberMode = config[0].NumberMode
		emitEmpty = config[0].EmitEmpty
	}

	return NewPartitioner(PartitionerConfig{
		Delimiters: delimiters,
		SplitCamel: true,
		NumberMode: numberMode,
		EmitEmpty:  emitEmpty,
	})
}

// PartsToWords converts Parts to Words using classification logic.
func PartsToWords(parts []Part, config *ParserConfig) []Word {
	var words []Word
	for _, part := range parts {
		words = append(words, ClassifyPart(part, config))
	}
	return words
}

// ClassifyPart converts a Part into a Word.
func ClassifyPart(part Part, config *ParserConfig) Word {
	s := part.String()
	// Separator handling: if we have SeparatorPart, we might return it as a SeparatorWord?
	if _, ok := part.(*SeparatorPart); ok {
		return SeparatorWord(s)
	}

	if s == "" {
		return ExactCaseWord("")
	}

	// Check for dots -> Acronym
	if strings.Contains(s, ".") {
		return AcronymWord(s)
	}

	// Check casing
	isAllUpper := true
	isAllLower := true
	isTitle := false

	runes := []rune(s)
	if len(runes) > 0 && unicode.IsUpper(runes[0]) {
		isTitle = true
	}

	for i, r := range runes {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			isAllUpper = false
		}
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			isAllLower = false
		}
		if i > 0 && unicode.IsUpper(r) {
			isTitle = false
		}
	}

	if isAllUpper {
		// Use SmartAcronyms config or default
		smartAcronyms := true
		if config != nil {
			smartAcronyms = config.SmartAcronyms
		}

		if smartAcronyms && len(runes) > 1 {
			return AcronymWord(s)
		}
		return UpperCaseWord(s)
	}

	if isAllLower {
		return SingleCaseWord(s)
	}

	if isTitle {
		return FirstUpperCaseWord(s)
	}

	return ExactCaseWord(s)
}

// Level 1 / 2 Helpers

func ParseSnakeCase(input string, opts ...any) ([]Word, error) {
	// Snake case implies split by underscore.
	combinedOpts := make([]any, 0, len(opts)+1)
	combinedOpts = append(combinedOpts, opts...)
	combinedOpts = append(combinedOpts, WithSnakeCasePartitioner())
	return Parse(input, combinedOpts...)
}

func ParseCamelCase(input string, opts ...any) ([]Word, error) {
	combinedOpts := make([]any, 0, len(opts)+1)
	combinedOpts = append(combinedOpts, opts...)
	combinedOpts = append(combinedOpts, WithCamelCasePartitioner())
	return Parse(input, combinedOpts...)
}

func ParseKebabCase(input string, opts ...any) ([]Word, error) {
	combinedOpts := make([]any, 0, len(opts)+1)
	combinedOpts = append(combinedOpts, opts...)
	combinedOpts = append(combinedOpts, WithKebabCasePartitioner())
	return Parse(input, combinedOpts...)
}
