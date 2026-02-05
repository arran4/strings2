package strings2

import (
	"strings"
	"unicode"
)

// Parser is a configurable string parser.
type Parser struct {
	SmartAcronyms bool

	// Delimiters is a list of characters that can be used as word separators.
	Delimiters []rune

	// SplitCamelCase enables splitting on case transitions (e.g. lower->Upper).
	SplitCamelCase bool
}

// ParserOption defines a function to configure the Parser.
type ParserOption func(*Parser)

// ParserSmartAcronyms enables intelligent acronym detection (e.g. treating all-caps words as AcronymWord).
func ParserSmartAcronyms(enabled bool) ParserOption {
	return func(p *Parser) {
		p.SmartAcronyms = enabled
	}
}

// ParserDelimiters sets the allowed delimiters for detection.
// Common delimiters are '_', '-', ' ', '.'.
func ParserDelimiters(delims ...rune) ParserOption {
	return func(p *Parser) {
		p.Delimiters = delims
	}
}

// ParserSplitCamelCase enables or disables splitting on casing transitions.
func ParserSplitCamelCase(enabled bool) ParserOption {
	return func(p *Parser) {
		p.SplitCamelCase = enabled
	}
}

// Parse parses the input string into a slice of Words based on detected or provided configuration.
func Parse(input string, opts ...ParserOption) ([]Word, error) {
	p := &Parser{
		SmartAcronyms:  true,
		SplitCamelCase: true,
		Delimiters:     []rune{'_', '-', ' '},
	}
	for _, opt := range opts {
		opt(p)
	}

	parts := p.split(input)

	words := make([]Word, len(parts))
	for i, part := range parts {
		words[i] = p.classify(part)
	}

	return words, nil
}

func (p *Parser) split(input string) []string {
	if input == "" {
		return []string{""}
	}

	// Heuristic: Find the delimiter with the highest occurrence count.
	var bestDelim rune
	maxCount := 0

	for _, d := range p.Delimiters {
		count := strings.Count(input, string(d))
		if count > maxCount {
			maxCount = count
			bestDelim = d
		}
	}

	// If a delimiter is found, use it.
	if maxCount > 0 {
		// Special handling for "Screaming Snake Case" or similar where purely splitting might need case adjustment?
		// But for now, just split.
		// Note: strings.Fields is special for space, but strings.Split is strict.
		// User mentioned "context", sticking to strict split by best delim for now.
		// If delimiter is space, use Fields to handle multiple spaces nicely?
		if bestDelim == ' ' {
			return strings.Fields(input)
		}
		return strings.Split(input, string(bestDelim))
	}

	// If no delimiters found, check for CamelCase if enabled.
	if p.SplitCamelCase {
		// Only try CamelCase if we see mixed casing?
		// Actually, splitCamelCase handles the logic.
		return splitCamelCase(input)
	}

	// Fallback: Return whole string.
	return []string{input}
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

	// Check for dots -> Acronym (unless dot was the delimiter, but split removes delimiters)
	if strings.Contains(part, ".") {
		return AcronymWord(part)
	}

	// Check casing
	isAllUpper := true
	isAllLower := true
	isTitle := false

	runes := []rune(part)
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
