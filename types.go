package strings2

import (
	"fmt"
	"strings"
)

// Word interface representing a stringer type
type Word fmt.Stringer

// Word Types
type SingleCaseWord string
type FirstUpperCaseWord string
type ExactCaseWord string

// String implementations
func (w SingleCaseWord) String() string     { return strings.ToLower(string(w)) }
func (w FirstUpperCaseWord) String() string { return UpperCaseFirst(strings.ToLower(string(w))) }

func UpperCaseFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
func (w ExactCaseWord) String() string { return string(w) }

// Options
type Option func(*caseConfig)

type CaseMode int

const (
	CMVerbatim CaseMode = iota
	CMFirstTitle
	CMAllTitle
	CMFirstLower
	CMWhispering
	CMScreaming
)

func (cm CaseMode) Transform(word []Word, wordPos int) []string {
	if word == nil || len(word) == 0 {
		return word
	}
	switch cm {
	case CMVerbatim:
		return word
	case CMFirstTitle:
		switch word[0].(type) {
		case SingleCaseWord:
			if wordPos == 0 {
				word[0] = UpperCaseFirst(word[0])
			}
		}
	}
}

type caseConfig struct {
	caseMode       CaseMode
	delimiter      string
	upperIndicator string
}

func OptionDelimiter(d string) Option {
	return func(cfg *caseConfig) { cfg.delimiter = d }
}

func OptionCaseMode(caseMode CaseMode) Option {
	return func(cfg *caseConfig) { cfg.caseMode = caseMode }
}

// ToFormattedCase generates formatted case strings with the given options
func ToFormattedCase(words []Word, opts ...Option) string {
	cfg := &caseConfig{delimiter: "-"}

	for _, opt := range opts {
		opt(cfg)
	}

	var result []string
	for _, word := range words {
		switch word := word.(type) {
		case SingleCaseWord:
			w := word.String()
			if cfg.allUpper || cfg.screaming {
				w = strings.ToUpper(w)
			} else if cfg.allLower || cfg.whispering {
				w = strings.ToLower(w)
			}
		case ExactCaseWord:
			if cfg.mixCaseSupport {
				w = splitMixCase(w, cfg.delimiter)
			}
		case FirstUpperCaseWord:
			if cfg.mixCaseSupport {
				w = splitMixCase(w, cfg.delimiter)
			}
		}

		result = append(result, w)
	}

	delimiter := cfg.delimiter
	if cfg.upperIndicator == cfg.delimiter {
		delimiter = cfg.delimiter + cfg.delimiter
	}
	final := strings.Join(result, delimiter)

	// Special handling for CamelCase
	if !cfg.firstUpper {
		final = strings.ToLower(final[:1]) + final[1:]
	}

	return final
}

// Helper function to split words in mixed case
func splitMixCase(input, delimiter string) string {
	var result strings.Builder
	for i, r := range input {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteString(delimiter)
		}
		result.WriteRune(r)
	}
	return result.String()
}

// ToKebabCase converts words into kebab-case format.
//
// Options:
//   - Delimiter: Defaults to "-".
//   - DoubleDelimiter: Uses a double "-" to signify reused delimiters.
func ToKebabCase(words []Word, opts ...Option) string {
	return ToFormattedCase(words, append([]Option(nil), append(opts, OptionDelimiter("-"))...)...)
}

// ToSnakeCase converts words into snake_case format.
//
// Options:
//   - Delimiter: Defaults to "_".
//   - Screaming: Converts the entire output to upper case (SCREAMING_SNAKE_CASE).
func ToSnakeCase(words []Word, opts ...Option) string {
	return ToFormattedCase(words, append([]Option(nil), append(opts, OptionDelimiter("_"))...)...)
}

// ToPascalCase converts words into PascalCase format.
//
// Options:
//   - FirstUpper: Ensures the first letter of the result is uppercase.
func ToPascalCase(words []Word, opts ...Option) string {
	return ToFormattedCase(words, append([]Option(nil), append(opts, OptionDelimiter(""), OptionFirstUpper())...)...)
}

// ToCamelCase converts words into camelCase format.
//
// Options:
//   - FirstUpper: Ensures the first letter of the result is uppercase (default is lowercase).
func ToCamelCase(words []Word, opts ...Option) string {
	return ToFormattedCase(words, append([]Option(nil), append(opts, OptionDelimiter(""), OptionFirstLower())...)...)
}
