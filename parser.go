package strings2

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Format defines the known string formats.
type Format int

const (
	FormatUnknown Format = iota
	FormatCamelCase
	FormatPascalCase
	FormatSnakeCase
	FormatKebabCase
	FormatScreamingSnakeCase
	FormatSentence
)

// Parser is a configurable string parser.
type Parser struct {
	SmartAcronyms bool

	// Feature flags
	CamelCase bool
	SnakeCase bool
	KebabCase bool
	Sentence  bool
}

// ParserOption defines a function to configure the Parser.
type ParserOption func(*Parser)

// ParserSmartAcronyms enables intelligent acronym detection (e.g. treating all-caps words as AcronymWord).
func ParserSmartAcronyms(enabled bool) ParserOption {
	return func(p *Parser) {
		p.SmartAcronyms = enabled
	}
}

// ParserCamelCase enables or disables CamelCase and PascalCase detection.
func ParserCamelCase(enabled bool) ParserOption {
	return func(p *Parser) {
		p.CamelCase = enabled
	}
}

// ParserSnakeCase enables or disables SnakeCase detection (including ScreamingSnakeCase).
func ParserSnakeCase(enabled bool) ParserOption {
	return func(p *Parser) {
		p.SnakeCase = enabled
	}
}

// ParserKebabCase enables or disables KebabCase detection.
func ParserKebabCase(enabled bool) ParserOption {
	return func(p *Parser) {
		p.KebabCase = enabled
	}
}

// ParserSentence enables or disables Sentence detection.
func ParserSentence(enabled bool) ParserOption {
	return func(p *Parser) {
		p.Sentence = enabled
	}
}

// Parse parses the input string into a slice of Words based on detected or provided configuration.
func Parse(input string, opts ...ParserOption) ([]Word, error) {
	p := &Parser{
		SmartAcronyms: true,
		CamelCase:     true,
		SnakeCase:     true,
		KebabCase:     true,
		Sentence:      true,
	}
	for _, opt := range opts {
		opt(p)
	}

	format := p.detectFormat(input)

	parts := p.split(input, format)

	words := make([]Word, len(parts))
	for i, part := range parts {
		words[i] = p.classify(part)
	}

	return words, nil
}

func (p *Parser) detectFormat(input string) Format {
	if p.SnakeCase && strings.Contains(input, "_") {
		if input == strings.ToUpper(input) {
			return FormatScreamingSnakeCase
		}
		return FormatSnakeCase
	}
	if p.KebabCase && strings.Contains(input, "-") {
		return FormatKebabCase
	}
	if p.Sentence && strings.Contains(input, " ") {
		return FormatSentence
	}
	if input == "" {
		return FormatUnknown
	}

	if p.CamelCase {
		firstRune, _ := utf8.DecodeRuneInString(input)
		if unicode.IsUpper(firstRune) {
			return FormatPascalCase
		}
		return FormatCamelCase
	}

	return FormatUnknown
}

func (p *Parser) split(input string, format Format) []string {
	switch format {
	case FormatSnakeCase, FormatScreamingSnakeCase:
		return strings.Split(input, "_")
	case FormatKebabCase:
		return strings.Split(input, "-")
	case FormatSentence:
		return strings.Fields(input)
	case FormatCamelCase, FormatPascalCase:
		return splitCamelCase(input)
	default:
		// Fallback for Unknown: Return whole string or try to split by non-alphanum?
		// Returning whole string as one part seems safest.
		return []string{input}
	}
}

func splitCamelCase(input string) []string {
	var parts []string
	runes := []rune(input)
	start := 0
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if i > 0 {
			prev := runes[i-1]
			var next rune
			if i+1 < len(runes) {
				next = runes[i+1]
			}

			// Case 1: lower -> Upper
			if unicode.IsLower(prev) && unicode.IsUpper(r) {
				parts = append(parts, string(runes[start:i]))
				start = i
			}

			// Case 2: Upper -> Upper -> lower (e.g. PDFLoader, split at L)
			if unicode.IsUpper(prev) && unicode.IsUpper(r) && unicode.IsLower(next) {
				parts = append(parts, string(runes[start:i]))
				start = i
			}
		}
	}
	parts = append(parts, string(runes[start:]))
	return parts
}

func (p *Parser) classify(part string) Word {
	if part == "" {
		return ExactCaseWord("")
	}

	// Check for dots -> Acronym
	if strings.Contains(part, ".") {
		return AcronymWord(part)
	}

	// Check casing
	isAllUpper := true
	isAllLower := true
	isTitle := false

	runes := []rune(part)
	if unicode.IsUpper(runes[0]) {
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
			isTitle = false // Title case has only first upper
		}
	}

	if isAllUpper {
		if p.SmartAcronyms && len(runes) > 1 {
			// Heuristic: Multi-letter all-upper is Acronym
			return AcronymWord(part)
		}
		return UpperCaseWord(part)
	}

	if isAllLower {
		return SingleCaseWord(part)
	}

	if isTitle {
		return FirstUpperCaseWord(part)
	}

	return ExactCaseWord(part)
}
