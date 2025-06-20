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

func (cm CaseMode) Transform(words []Word, pos int) []Word {
	if len(words) == 0 || pos >= len(words) {
		return words
	}

	switch cm {
	case CMFirstTitle:
		if pos == 0 {
			words[pos] = SingleCaseWord(UpperCaseFirst(words[pos].String()))
		}
	case CMAllTitle:
		words[pos] = SingleCaseWord(UpperCaseFirst(words[pos].String()))
	case CMFirstLower:
		if pos == 0 {
			words[pos] = SingleCaseWord(strings.ToLower(words[pos].String()))
		}
	case CMWhispering:
		words[pos] = SingleCaseWord(strings.ToLower(words[pos].String()))
	case CMScreaming:
		words[pos] = SingleCaseWord(strings.ToUpper(words[pos].String()))
	}

	return words
}

type caseConfig struct {
	caseMode       CaseMode
	delimiter      string
	upperIndicator string
	firstUpper     bool
}

func OptionDelimiter(d string) Option {
	return func(cfg *caseConfig) { cfg.delimiter = d }
}

func OptionCaseMode(caseMode CaseMode) Option {
	return func(cfg *caseConfig) { cfg.caseMode = caseMode }
}

// OptionFirstUpper ensures the resulting string begins with an upper case letter.
func OptionFirstUpper() Option {
	return func(cfg *caseConfig) { cfg.firstUpper = true }
}

// OptionFirstLower ensures the resulting string begins with a lower case letter.
func OptionFirstLower() Option {
	return func(cfg *caseConfig) { cfg.firstUpper = false }
}

// ToFormattedCase generates formatted case strings with the given options
func ToFormattedCase(words []Word, opts ...Option) string {
	cfg := &caseConfig{delimiter: "-"}

	for _, opt := range opts {
		opt(cfg)
	}

	var result []string
	for i, w := range words {
		out := w.String()

		switch cfg.caseMode {
		case CMFirstTitle:
			if i == 0 {
				out = UpperCaseFirst(strings.ToLower(out))
			} else {
				out = strings.ToLower(out)
			}
		case CMAllTitle:
			out = UpperCaseFirst(strings.ToLower(out))
		case CMFirstLower:
			if i == 0 {
				out = strings.ToLower(out)
			}
		case CMWhispering:
			out = strings.ToLower(out)
		case CMScreaming:
			out = strings.ToUpper(out)
		}

		result = append(result, out)
	}

	delimiter := cfg.delimiter
	if cfg.upperIndicator == cfg.delimiter {
		delimiter = cfg.delimiter + cfg.delimiter
	}
	final := strings.Join(result, delimiter)

	if delimiter == "" && !cfg.firstUpper && len(final) > 0 {
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
