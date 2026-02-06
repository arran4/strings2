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
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case ParserOption:
			o.Apply(p)
		case Partitioner:
			p.Partitioner = o
		case ParserSmartAcronyms:
			o.Apply(p)
		}

		// Also check if it implements ParserOption directly (if not caught by switch case above, though ParserOption case should cover it if interface match)
		// But in Go, type switch cases on interface type will match if concrete type implements it.
	}

	// Level 4: Partition
	// If partitioner is not set, try to detect
	partitioner := p.Partitioner
	if partitioner == nil {
		partitioner = DetectPartitioner(stats)
	}

	parts := SubPartsToParts(subs, partitioner)

	// Level 3: Words
	words := PartsToWords(parts, p)

	return words, nil
}

// ParserConfig holds configuration for the parsing pipeline.
type ParserConfig struct {
	Partitioner   Partitioner
	// SmartAcronyms controls whether all-uppercase words (longer than 1 char)
	// should be treated as AcronymWord instead of UpperCaseWord.
	// Defaults to true.
	SmartAcronyms bool
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

// DetectPartitioner uses stats to guess the best partitioner.
func DetectPartitioner(stats Stats) Partitioner {
	// Heuristic:
	// If spaces > 0, likely Sentence (Sentence usually beats Kebab/Snake if mixed)
	if stats.Spaces > 0 {
		return func(subs []SubPart) []Part {
			// Space partitioner
			var parts []Part
			var current []SubPart
			for _, s := range subs {
				if s.IsSpace() {
					if len(current) > 0 {
						parts = append(parts, &WordPart{BasePart{Subs: current}})
						current = nil
					}
				} else {
					current = append(current, s)
				}
			}
			if len(current) > 0 {
				parts = append(parts, &WordPart{BasePart{Subs: current}})
			}
			return parts
		}
	}
	// If underscores > 0, likely SnakeCase
	if stats.SymbolCounts['_'] > 0 {
		return SnakeCasePartitioner
	}
	// If hyphens > 0, likely KebabCase
	if stats.SymbolCounts['-'] > 0 {
		return KebabCasePartitioner
	}

	// Default to CamelCase
	return CamelCasePartitioner
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
