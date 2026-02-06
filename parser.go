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
// - ParserSmartAcronyms bool
func Parse(input string, opts ...any) ([]Word, error) {
	// Level 5: Scan
	subs, stats := StringToSubParts(input)

	p := &ParserConfig{
		SmartAcronyms: true,
		NumberSplitting: false, // Default to false to preserve http200
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case Partitioner:
			p.Partitioner = o
		case ParserOption:
			o.Apply(p)
		}
	}

	// Level 4: Partition
	// If partitioner is not set, try to detect
	partitioner := p.Partitioner
	if partitioner == nil {
		partitioner = DetectPartitioner(stats, p)
	}

	parts := SubPartsToParts(subs, partitioner)

	// Level 3: Words
	words := PartsToWords(parts, p)

	return words, nil
}

// ParserConfig holds configuration for the parsing pipeline.
type ParserConfig struct {
	Partitioner Partitioner
	// SmartAcronyms controls whether all-uppercase words (longer than 1 char)
	// should be treated as AcronymWord instead of UpperCaseWord.
	// Defaults to true.
	SmartAcronyms bool
	// NumberSplitting controls whether to split on letter-digit boundaries.
	NumberSplitting bool
}

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
func WithNumberSplitting(enabled bool) ParserOption {
	return funcParserOption(func(p *ParserConfig) {
		p.NumberSplitting = enabled
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

	splitNumber := false
	if len(config) > 0 && config[0] != nil {
		splitNumber = config[0].NumberSplitting
	}

	return NewPartitioner(delimiters, true, splitNumber)
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

func ParseSnakeCase(input string) []Word {
	subs, _ := StringToSubParts(input)
	parts := SubPartsToParts(subs, SnakeCasePartitioner)
	return PartsToWords(parts, nil)
}

func ParseCamelCase(input string) []Word {
	subs, _ := StringToSubParts(input)
	parts := SubPartsToParts(subs, CamelCasePartitioner)
	return PartsToWords(parts, nil)
}

func ParseKebabCase(input string) []Word {
	subs, _ := StringToSubParts(input)
	parts := SubPartsToParts(subs, KebabCasePartitioner)
	return PartsToWords(parts, nil)
}
