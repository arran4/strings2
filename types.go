package strings2

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrRune = errors.New("invalid rune")

// Word interface representing a stringer type that can be used in casing conversions.
type Word fmt.Stringer

// Word Types

// SingleCaseWord is a word that will be lowercased when stringified.
type SingleCaseWord string

// FirstUpperCaseWord is a word that will have its first letter uppercased and the rest lowercased when stringified.
type FirstUpperCaseWord string

// ExactCaseWord is a word that preserves its case when stringified.
type ExactCaseWord string

// AcronymWord is a word that represents an acronym.
// It is usually preserved in case, but can be configured otherwise.
type AcronymWord string

// UpperCaseWord is a word that was originally all uppercase.
type UpperCaseWord string

// SeparatorWord is a delimiter or separator preserved from the input.
type SeparatorWord string

// String implementations
func (w SingleCaseWord) String() string     { return strings.ToLower(string(w)) }
func (w FirstUpperCaseWord) String() string { return upperCaseFirstLower(string(w)) }
func (w AcronymWord) String() string        { return string(w) }
func (w UpperCaseWord) String() string      { return strings.ToUpper(string(w)) }
func (w SeparatorWord) String() string      { return string(w) }

func performCaseFirst(s string, fn func(rune) rune) (string, rune, bool) {
	if s == "" {
		return s, 0, true
	}
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return "", r, false
	}
	u := fn(r)
	if r == u {
		return s, 0, true
	}
	var b strings.Builder
	b.Grow(len(s) + utf8.UTFMax)
	b.WriteRune(u)
	b.WriteString(s[size:])
	return b.String(), 0, true
}

// UpperCaseFirst uppercases the first character of the string.
func UpperCaseFirst(s string) string {
	res, _, ok := performCaseFirst(s, unicode.ToUpper)
	if !ok {
		return s
	}
	return res
}

// UpperCaseFirstWithErr uppercases the first character of the string.
// It returns an error if the first character is an invalid rune.
func UpperCaseFirstWithErr(s string) (string, error) {
	res, r, ok := performCaseFirst(s, unicode.ToUpper)
	if !ok {
		return "", fmt.Errorf("%w: %q", ErrRune, r)
	}
	return res, nil
}

// MustUpperCaseFirst uppercases the first character of the string.
// It panics if the first character is an invalid rune.
func MustUpperCaseFirst(s string) string {
	res, _, ok := performCaseFirst(s, unicode.ToUpper)
	if !ok {
		panic(s)
	}
	return res
}

// LowerCaseFirst lowercases the first character of the string.
func LowerCaseFirst(s string) string {
	res, _, ok := performCaseFirst(s, unicode.ToLower)
	if !ok {
		return s
	}
	return res
}

// LowerCaseFirstWithErr lowercases the first character of the string.
// It returns an error if the first character is an invalid rune.
func LowerCaseFirstWithErr(s string) (string, error) {
	res, r, ok := performCaseFirst(s, unicode.ToLower)
	if !ok {
		return "", fmt.Errorf("%w: %q", ErrRune, r)
	}
	return res, nil
}

// MustLowerCaseFirst lowercases the first character of the string.
// It panics if the first character is an invalid rune.
func MustLowerCaseFirst(s string) string {
	res, _, ok := performCaseFirst(s, unicode.ToLower)
	if !ok {
		panic(s)
	}
	return res
}

// upperCaseFirstLower capitalizes the first character and lowercases the rest.
func upperCaseFirstLower(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size == 1 {
		// Invalid UTF-8 start byte.
		// We want to replace it with RuneError (like strings.ToLower/ToUpper do).
		// So we force needChange.
	} else if r == utf8.RuneError {
		// Valid RuneError (U+FFFD)
	}

	u := unicode.ToUpper(r)

	// Check if changes are needed
	needChange := (r != u) || (r == utf8.RuneError && size == 1)
	if !needChange {
		for _, rc := range s[size:] {
			if unicode.ToLower(rc) != rc {
				needChange = true
				break
			}
		}
	}

	if !needChange {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))
	b.WriteRune(u)
	for _, rc := range s[size:] {
		b.WriteRune(unicode.ToLower(rc))
	}
	return b.String()
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
// Deprecated: Use WordsToFormattedCase. This function suppresses errors for backward compatibility.
func ToFormattedCase(words []Word, opts ...Option) string {
	res, _ := WordsToFormattedCase(words, convertOptions(opts)...)
	return res
}

func convertOptions(opts []Option) []any {
	var out []any
	for _, o := range opts {
		out = append(out, o)
	}
	return out
}

// WordsToFormattedCase generates formatted case strings with the given options
func WordsToFormattedCase(words []Word, opts ...any) (string, error) {
	cfg := &caseConfig{delimiter: "-"}

	for _, opt := range opts {
		if o, ok := opt.(Option); ok {
			o(cfg)
		}
	}

	if cfg.upperIndicator != "" {
		if cfg.upperIndicator == cfg.delimiter {
			cfg.delimiter = cfg.delimiter + cfg.delimiter
		} else {
			cfg.delimiter = cfg.upperIndicator
		}
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

	result := make([]string, 0, len(words))
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
				w = upperCaseFirstLower(w)
			} else {
				w = strings.ToLower(w)
			}
		case ExactCaseWord:
			w = word.String()
			if cfg.mixCaseSupport {
				w = splitMixCase(w, cfg.delimiter)
			}
			if cfg.allUpper || cfg.screaming {
				w = strings.ToUpper(w)
			} else if cfg.allLower || cfg.whispering {
				w = strings.ToLower(w)
			}
		case FirstUpperCaseWord:
			w = word.String()
			if cfg.mixCaseSupport {
				w = splitMixCase(w, cfg.delimiter)
			}
			if cfg.allUpper || cfg.screaming {
				w = strings.ToUpper(w)
			} else if cfg.allLower || cfg.whispering {
				w = strings.ToLower(w)
			}
		case AcronymWord:
			w = word.String()
			if cfg.screaming {
				w = strings.ToUpper(w)
			} else if cfg.whispering {
				w = strings.ToLower(w)
			} else if cfg.caseMode == CMAllTitle {
				w = upperCaseFirstLower(w)
			}
		case UpperCaseWord:
			w = word.String()
			if cfg.allUpper || cfg.screaming {
				w = strings.ToUpper(w)
			} else if cfg.allLower || cfg.whispering {
				w = strings.ToLower(w)
			} else if cfg.caseMode == CMAllTitle {
				w = upperCaseFirstLower(w)
			} else {
				w = strings.ToLower(w)
			}
		case SeparatorWord:
			w = word.String()
		default:
			w = word.String()
		}

		result = append(result, w)
	}

	final := strings.Join(result, cfg.delimiter)

	if cfg.firstUpper {
		final = UpperCaseFirst(final)
	}
	if cfg.firstLower {
		final = LowerCaseFirst(final)
	}

	return final, nil
}

// PartsToFormattedCase converts Parts to words then formats them.
// This is useful when you have intermediate Parts and want to format them directly.
func PartsToFormattedCase(parts []Part, opts ...any) (string, error) {
	// Extract ParserConfig from opts to use for classification
	p := &ParserConfig{
		SmartAcronyms:   true,
		NumberSplitting: false,
	}
	for _, opt := range opts {
		if o, ok := opt.(ParserOption); ok {
			o.Apply(p)
		}
	}

	words := PartsToWords(parts, p)
	return WordsToFormattedCase(words, opts...)
}

// ToFormattedString converts string to formatted case (generic entry point).
func ToFormattedString(s string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptionsAny(opts)
	words, err := Parse(s, parseOpts...)
	if err != nil {
		return "", err
	}
	return WordsToFormattedCase(words, fmtOpts...)
}

// FromFormattedString is alias for ToFormattedString.
func FromFormattedString(s string, opts ...any) (string, error) {
	return ToFormattedString(s, opts...)
}

// Helper to split ...any into parse opts and format opts (returning as []any for flexibility)
func separateOptionsAny(opts []any) ([]any, []any) {
	var parseOpts []any
	var fmtOpts []any

	for _, o := range opts {
		switch v := o.(type) {
		case Option:
			fmtOpts = append(fmtOpts, v)
		case ParserOption, Partitioner, PartitionerConfig:
			parseOpts = append(parseOpts, v)
		default:
			// Assume unknown types might be relevant for formatter if it changes,
			// or just ignore.
		}
	}
	return parseOpts, fmtOpts
}

// Helper function to split words in mixed case
func splitMixCase(input, delimiter string) string {
	if delimiter == "" {
		return input
	}
	var result strings.Builder
	// Pre-allocate to avoid resizing.
	// We add a buffer for potential delimiters (assuming roughly 50% increase).
	result.Grow(len(input) + len(input)/2)
	for i, r := range input {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteString(delimiter)
		}
		result.WriteRune(r)
	}
	return result.String()
}

// ToKebabCase converts words into kebab-case format.
func ToKebabCase(words []Word, opts ...Option) (string, error) {
	return WordsToFormattedCase(words, append(convertOptions(opts), OptionDelimiter("-"))...)
}

// ToSnakeCase converts words into snake_case format.
func ToSnakeCase(words []Word, opts ...Option) (string, error) {
	return WordsToFormattedCase(words, append(convertOptions(opts), OptionDelimiter("_"))...)
}

// ToPascalCase converts words into PascalCase format.
func ToPascalCase(words []Word, opts ...Option) (string, error) {
	return WordsToFormattedCase(words, append(convertOptions(opts), OptionDelimiter(""), OptionFirstUpper(), OptionCaseMode(CMAllTitle))...)
}

// ToCamelCase converts words into camelCase format.
func ToCamelCase(words []Word, opts ...Option) (string, error) {
	return WordsToFormattedCase(words, append(convertOptions(opts), OptionDelimiter(""), OptionFirstLower(), OptionCaseMode(CMAllTitle))...)
}
