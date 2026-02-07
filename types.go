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
func (w FirstUpperCaseWord) String() string { return UpperCaseFirst(strings.ToLower(string(w))) }
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
	return string(u) + s[size:], 0, true
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

	delimiter := cfg.delimiter
	if cfg.upperIndicator != "" {
		if cfg.upperIndicator == cfg.delimiter {
			delimiter = cfg.delimiter + cfg.delimiter
		} else {
			delimiter = cfg.upperIndicator
		}
	}

	size := 0
	for _, word := range words {
		switch w := word.(type) {
		case SingleCaseWord:
			size += len(w)
		case ExactCaseWord:
			size += len(w)
		case FirstUpperCaseWord:
			size += len(w)
		case AcronymWord:
			size += len(w)
		case UpperCaseWord:
			size += len(w)
		case SeparatorWord:
			size += len(w)
		default:
			size += 5 // fallback
		}
	}
	size += len(delimiter) * max(0, len(words)-1)

	var b strings.Builder
	b.Grow(size)

	for i, word := range words {
		if i > 0 {
			b.WriteString(delimiter)
		}

		switch word := word.(type) {
		case SingleCaseWord:
			s := string(word)
			if cfg.allUpper || cfg.screaming {
				for _, r := range s {
					b.WriteRune(unicode.ToUpper(r))
				}
			} else if cfg.allLower || cfg.whispering {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			} else if cfg.caseMode == CMAllTitle {
				first := true
				for _, r := range s {
					if first {
						b.WriteRune(unicode.ToUpper(r))
						first = false
					} else {
						b.WriteRune(unicode.ToLower(r))
					}
				}
			} else {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			}
		case ExactCaseWord:
			s := string(word)
			if cfg.mixCaseSupport {
				for j, r := range s {
					if j > 0 && unicode.IsUpper(r) {
						if cfg.allUpper || cfg.screaming {
							for _, dr := range cfg.delimiter {
								b.WriteRune(unicode.ToUpper(dr))
							}
						} else if cfg.allLower || cfg.whispering {
							for _, dr := range cfg.delimiter {
								b.WriteRune(unicode.ToLower(dr))
							}
						} else {
							b.WriteString(cfg.delimiter)
						}
					}
					if cfg.allUpper || cfg.screaming {
						b.WriteRune(unicode.ToUpper(r))
					} else if cfg.allLower || cfg.whispering {
						b.WriteRune(unicode.ToLower(r))
					} else {
						b.WriteRune(r)
					}
				}
			} else {
				if cfg.allUpper || cfg.screaming {
					for _, r := range s {
						b.WriteRune(unicode.ToUpper(r))
					}
				} else if cfg.allLower || cfg.whispering {
					for _, r := range s {
						b.WriteRune(unicode.ToLower(r))
					}
				} else {
					b.WriteString(s)
				}
			}
		case FirstUpperCaseWord:
			s := string(word)
			if cfg.allUpper || cfg.screaming {
				for _, r := range s {
					b.WriteRune(unicode.ToUpper(r))
				}
			} else if cfg.allLower || cfg.whispering {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			} else {
				first := true
				for _, r := range s {
					if first {
						b.WriteRune(unicode.ToUpper(r))
						first = false
					} else {
						b.WriteRune(unicode.ToLower(r))
					}
				}
			}
		case AcronymWord:
			s := string(word)
			if cfg.screaming {
				for _, r := range s {
					b.WriteRune(unicode.ToUpper(r))
				}
			} else if cfg.whispering {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			} else if cfg.caseMode == CMAllTitle {
				first := true
				for _, r := range s {
					if first {
						b.WriteRune(unicode.ToUpper(r))
						first = false
					} else {
						b.WriteRune(unicode.ToLower(r))
					}
				}
			} else {
				b.WriteString(s)
			}
		case UpperCaseWord:
			s := string(word)
			if cfg.allUpper || cfg.screaming {
				for _, r := range s {
					b.WriteRune(unicode.ToUpper(r))
				}
			} else if cfg.allLower || cfg.whispering {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			} else if cfg.caseMode == CMAllTitle {
				first := true
				for _, r := range s {
					if first {
						b.WriteRune(unicode.ToUpper(r))
						first = false
					} else {
						b.WriteRune(unicode.ToLower(r))
					}
				}
			} else {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			}
		case SeparatorWord:
			b.WriteString(string(word))
		default:
			b.WriteString(word.String())
		}
	}

	final := b.String()

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
