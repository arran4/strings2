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

	// IsDelimiter is a function that checks if a rune is a delimiter.
	IsDelimiter func(r rune) bool

	// IsNewSubPart is a function that determines if a new sub-part starts at the given position.
	IsNewSubPart func(pos int, surrounding []rune) bool
}

// ParserOption defines a function to configure the Parser.
type ParserOption interface {
	Apply(*Parser)
}

type funcParserOption func(*Parser)

func (f funcParserOption) Apply(p *Parser) {
	f(p)
}

// ParserSmartAcronyms enables intelligent acronym detection (e.g. treating all-caps words as AcronymWord).
type ParserSmartAcronyms bool

func (b ParserSmartAcronyms) Apply(p *Parser) {
	p.SmartAcronyms = bool(b)
}

// ParserDelimiters sets the allowed delimiters for detection.
type ParserDelimiters []rune

func (d ParserDelimiters) Apply(p *Parser) {
	p.Delimiters = d
}

// ParserSplitCamelCase enables or disables splitting on casing transitions.
type ParserSplitCamelCase bool

func (b ParserSplitCamelCase) Apply(p *Parser) {
	p.SplitCamelCase = bool(b)
}

// ParserIsDelimiter sets a custom function to check if a rune is a delimiter.
type ParserIsDelimiter func(r rune) bool

func (f ParserIsDelimiter) Apply(p *Parser) {
	p.IsDelimiter = f
}

// ParserIsNewSubPart sets a custom function to determine sub-part boundaries.
type ParserIsNewSubPart func(pos int, surrounding []rune) bool

func (f ParserIsNewSubPart) Apply(p *Parser) {
	p.IsNewSubPart = f
}

// Parse parses the input string into a slice of Words based on detected or provided configuration.
func Parse(input string, opts ...ParserOption) ([]Word, error) {
	p := &Parser{
		SmartAcronyms:  true,
		SplitCamelCase: true,
		Delimiters:     []rune{'_', '-', ' '},
	}

	for _, opt := range opts {
		opt.Apply(p)
	}

	// Ensure IsDelimiter is set if Delimiters is used (default behavior)
	if p.IsDelimiter == nil && len(p.Delimiters) > 0 {
		// Heuristic: Find the primary delimiter (highest frequency in input)
		// If multiple delimiters exist, we only want to split by the primary one (or ones that act as delimiters in this context).
		// The test "Contextual Dash" implies that if Space is the primary delimiter, Dash should NOT be treated as a delimiter unless it's also high frequency?
		// Actually, the test "Hello to all the good-doers out there" expects "good-doers" to remain intact.
		// This means we should detect the DOMINANT delimiter and only use that.

		primaryDelim := detectPrimaryDelimiter(input, p.Delimiters)

		p.IsDelimiter = func(r rune) bool {
			return r == primaryDelim
		}
	}

	// Default CamelCase splitter if not provided but enabled
	if p.IsNewSubPart == nil && p.SplitCamelCase {
		p.IsNewSubPart = func(pos int, runes []rune) bool {
			if pos == 0 {
				return false
			}
			r := runes[pos]
			prev := runes[pos-1]

			// Case 1: lower -> Upper
			if unicode.IsLower(prev) && unicode.IsUpper(r) {
				return true
			}

			// Case 2: Upper -> Upper -> lower (e.g. PDFLoader, split at L)
			if pos+1 < len(runes) {
				next := runes[pos+1]
				if unicode.IsUpper(prev) && unicode.IsUpper(r) && unicode.IsLower(next) {
					return true
				}
			}
			return false
		}
	}

	parts := p.split(input)

	words := make([]Word, len(parts))
	for i, part := range parts {
		words[i] = p.classify(part)
	}

	return words, nil
}

func detectPrimaryDelimiter(input string, delimiters []rune) rune {
	maxCount := 0
	var primary rune
	for _, d := range delimiters {
		count := strings.Count(input, string(d))
		if count > maxCount {
			maxCount = count
			primary = d
		}
	}
	// If no delimiters found, just return 0 (null char), IsDelimiter will return false.
	return primary
}

func (p *Parser) split(input string) []string {
	if input == "" {
		return []string{""}
	}

	runes := []rune(input)
	var parts []string
	start := 0

	for i := 0; i < len(runes); i++ {
		// Check for delimiter split
		if p.IsDelimiter != nil && p.IsDelimiter(runes[i]) {
			if i > start {
				parts = append(parts, string(runes[start:i]))
			}
			start = i + 1
			continue
		}

		// Check for sub-part split (CamelCase)
		if p.IsNewSubPart != nil && p.IsNewSubPart(i, runes) {
			if i > start {
				parts = append(parts, string(runes[start:i]))
			}
			start = i
		}
	}

	if start < len(runes) {
		parts = append(parts, string(runes[start:]))
	} else if start == len(runes) && (p.IsDelimiter != nil && p.IsDelimiter(runes[len(runes)-1])) {
         // If ends with delimiter, ignoring empty trailing part usually desired
    }

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
