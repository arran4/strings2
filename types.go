package strings2

import (
	"fmt"
	"strings"
	"unicode"
)

// Word interface representing a stringer type that can be used in casing conversions.
type Word fmt.Stringer

// Word Types

// SingleCaseWord is a word that will be lowercased when stringified.
type SingleCaseWord string

// FirstUpperCaseWord is a word that will have its first letter uppercased and the rest lowercased when stringified.
type FirstUpperCaseWord string

// ExactCaseWord is a word that preserves its case when stringified.
type ExactCaseWord string

// String implementations
func (w SingleCaseWord) String() string     { return strings.ToLower(string(w)) }
func (w FirstUpperCaseWord) String() string { return UpperCaseFirst(strings.ToLower(string(w))) }

// UpperCaseFirst uppercases the first character of the string.
func UpperCaseFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
func (w ExactCaseWord) String() string { return string(w) }

// Options
type Option func(*caseConfig)

// CaseMode defines the casing transformation mode.
type CaseMode int

const (
	// CMVerbatim leaves the case as is.
	CMVerbatim CaseMode = iota
	// CMFirstTitle uppercases the first character of the first word.
	CMFirstTitle
	// CMAllTitle uppercases the first character of every word.
	CMAllTitle
	// CMFirstLower lowercases the first character of the first word.
	CMFirstLower
	// CMWhispering lowercases all characters (like snake_case or kebab-case usually).
	CMWhispering
	// CMScreaming uppercases all characters (like SCREAMING_SNAKE_CASE).
	CMScreaming
)

type caseConfig struct {
	caseMode       CaseMode
	delimiter      string
	upperIndicator string
	allUpper       bool
	allLower       bool
	screaming      bool
	whispering     bool
	mixCaseSupport bool
	firstUpper     bool
	firstLower     bool
}

// OptionDelimiter sets the delimiter between words.
func OptionDelimiter(d string) Option {
	return func(cfg *caseConfig) { cfg.delimiter = d }
}

// OptionCaseMode sets the case mode.
func OptionCaseMode(caseMode CaseMode) Option {
	return func(cfg *caseConfig) { cfg.caseMode = caseMode }
}

// OptionFirstUpper ensures the very first character of the result is uppercase.
func OptionFirstUpper() Option {
	return func(cfg *caseConfig) { cfg.firstUpper = true }
}

// OptionFirstLower ensures the very first character of the result is lowercase.
func OptionFirstLower() Option {
	return func(cfg *caseConfig) { cfg.firstLower = true }
}

// OptionMixCaseSupport enables splitting of mixed case words (e.g. CamelCase) into separate words based on uppercase letters.
func OptionMixCaseSupport() Option {
	return func(cfg *caseConfig) { cfg.mixCaseSupport = true }
}

// OptionUpperIndicator sets a specific indicator for upper case (often used for double delimiters).
func OptionUpperIndicator(d string) Option {
	return func(cfg *caseConfig) { cfg.upperIndicator = d }
}

// ToFormattedCase generates formatted case strings with the given options
func ToFormattedCase(words []Word, opts ...Option) string {
	cfg := &caseConfig{delimiter: "-"}

	for _, opt := range opts {
		opt(cfg)
	}

	switch cfg.caseMode {
	case CMScreaming:
		cfg.screaming = true
	case CMWhispering:
		cfg.whispering = true
	case CMFirstLower:
		cfg.firstLower = true
	case CMFirstTitle:
		cfg.firstUpper = true
	}

	var result []string
	for _, word := range words {
		var w string
		switch word := word.(type) {
		case SingleCaseWord:
			w = string(word)
			if cfg.allUpper || cfg.screaming {
				w = strings.ToUpper(w)
			} else if cfg.allLower || cfg.whispering {
				w = strings.ToLower(w)
			} else if cfg.caseMode == CMAllTitle {
				w = UpperCaseFirst(strings.ToLower(w))
			} else {
				w = strings.ToLower(w)
			}
		case ExactCaseWord:
			w = word.String()
			if cfg.mixCaseSupport {
				w = splitMixCase(w, cfg.delimiter)
			}
		case FirstUpperCaseWord:
			w = word.String()
			if cfg.mixCaseSupport {
				w = splitMixCase(w, cfg.delimiter)
			}
		default:
			w = word.String()
		}

		result = append(result, w)
	}

	delimiter := cfg.delimiter
	if cfg.upperIndicator != "" {
		if cfg.upperIndicator == cfg.delimiter {
			delimiter = cfg.delimiter + cfg.delimiter
		} else {
			delimiter = cfg.upperIndicator
		}
	}
	final := strings.Join(result, delimiter)

	if cfg.firstUpper {
		final = UpperCaseFirst(final)
	}
	if cfg.firstLower {
		if len(final) > 0 {
			final = strings.ToLower(final[:1]) + final[1:]
		}
	}

	return final
}

// Helper function to split words in mixed case
func splitMixCase(input, delimiter string) string {
	var result strings.Builder
	result.Grow(len(input))
	for i, r := range input {
		if i > 0 && unicode.IsUpper(r) {
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
	newOpts := make([]Option, 0, len(opts)+1)
	newOpts = append(newOpts, opts...)
	newOpts = append(newOpts, OptionDelimiter("-"))
	return ToFormattedCase(words, newOpts...)
}

// ToSnakeCase converts words into snake_case format.
//
// Options:
//   - Delimiter: Defaults to "_".
//   - Screaming: Converts the entire output to upper case (SCREAMING_SNAKE_CASE).
func ToSnakeCase(words []Word, opts ...Option) string {
	newOpts := make([]Option, 0, len(opts)+1)
	newOpts = append(newOpts, opts...)
	newOpts = append(newOpts, OptionDelimiter("_"))
	return ToFormattedCase(words, newOpts...)
}

// ToPascalCase converts words into PascalCase format.
//
// Options:
//   - FirstUpper: Ensures the first letter of the result is uppercase.
func ToPascalCase(words []Word, opts ...Option) string {
	newOpts := make([]Option, 0, len(opts)+3)
	newOpts = append(newOpts, opts...)
	newOpts = append(newOpts, OptionDelimiter(""), OptionFirstUpper(), OptionCaseMode(CMAllTitle))
	return ToFormattedCase(words, newOpts...)
}

// ToCamelCase converts words into camelCase format.
//
// Options:
//   - FirstUpper: Ensures the first letter of the result is uppercase (default is lowercase).
func ToCamelCase(words []Word, opts ...Option) string {
	newOpts := make([]Option, 0, len(opts)+3)
	newOpts = append(newOpts, opts...)
	newOpts = append(newOpts, OptionDelimiter(""), OptionFirstLower(), OptionCaseMode(CMAllTitle))
	return ToFormattedCase(words, newOpts...)
}
