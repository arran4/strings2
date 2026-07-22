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
func (w SingleCaseWord) String() string { return strings.ToLower(string(w)) }
func (w FirstUpperCaseWord) String() string {
	res, _ := upperCaseFirstLower(string(w), UTF8Replace)
	return res
}
func (w AcronymWord) String() string   { return string(w) }
func (w UpperCaseWord) String() string { return strings.ToUpper(string(w)) }
func (w SeparatorWord) String() string { return string(w) }

// Len implementations
func (w SingleCaseWord) Len() int     { return len(w) }
func (w FirstUpperCaseWord) Len() int { return len(w) }
func (w ExactCaseWord) Len() int      { return len(w) }
func (w AcronymWord) Len() int        { return len(w) }
func (w UpperCaseWord) Len() int      { return len(w) }
func (w SeparatorWord) Len() int      { return len(w) }


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
func upperCaseFirstLower(s string, mode UTF8Mode) (string, error) {
	if s == "" {
		return "", nil
	}
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size == 1 {
		if mode == UTF8Strict {
			return "", fmt.Errorf("%w: invalid rune", ErrRune)
		}
	}

	u := unicode.ToUpper(r)

	// Check if changes are needed.
	// If r == utf8.RuneError && size == 1, it is an invalid UTF-8 start byte.
	// We want to replace it with RuneError (like strings.ToLower/ToUpper do).
	// So we force needChange.
	needChange := (r != u) || (r == utf8.RuneError && size == 1 && mode == UTF8Replace)
	if !needChange {
		for _, rc := range s[size:] {
			if rc == utf8.RuneError {
				if mode == UTF8Strict {
					return "", fmt.Errorf("%w: invalid rune", ErrRune)
				}
			}
			if unicode.ToLower(rc) != rc {
				needChange = true
				break
			}
		}
	}

	if !needChange {
		return s, nil
	}

	var b strings.Builder
	b.Grow(len(s))
	if r == utf8.RuneError && size == 1 && mode == UTF8Ignore {
		b.WriteByte(s[0])
	} else {
		b.WriteRune(u)
	}

	for i, rc := range s[size:] {
		if rc == utf8.RuneError {
			if mode == UTF8Strict {
				return "", fmt.Errorf("%w: invalid rune", ErrRune)
			}
			if mode == UTF8Ignore {
				// s[size:] is the substring starting after first rune.
				// i is the index within that substring.
				// We need to write the original byte.
				// s[size+i] is the byte.
				b.WriteByte(s[size+i])
				continue
			}
		}
		b.WriteRune(unicode.ToLower(rc))
	}
	return b.String(), nil
}

func (w ExactCaseWord) String() string { return string(w) }

// WordLength returns the string length of the given Word type without allocating.
func WordLength(word Word) (int, error) {
	switch w := word.(type) {
	case SingleCaseWord:
		return len(w), nil
	case LowercaseWord:
		return len(w), nil
	case FirstUpperCaseWord:
		return len(w), nil
	case ExactCaseWord:
		return len(w), nil
	case AcronymWord:
		return len(w), nil
	case UpperCaseWord:
		return len(w), nil
	case SeparatorWord:
		return len(w), nil
	default:
		return 0, fmt.Errorf("unknown word type: %T", word)
	}
}

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
	// CMSmartTitle applies smart title case handling prepositions correctly.
	CMSmartTitle
)

// UTF8Mode defines how to handle invalid UTF-8 sequences.
type UTF8Mode int

const (
	// UTF8Replace replaces invalid UTF-8 bytes with utf8.RuneError (U+FFFD).
	UTF8Replace UTF8Mode = iota
	// UTF8Strict returns an error on invalid UTF-8 sequences.
	UTF8Strict
	// UTF8Ignore ignores invalid UTF-8 sequences and preserves the original bytes (best effort).
	UTF8Ignore
)

// FirstLowerBehavior determines how the first character of the formatted string should be lowercased.
type FirstLowerBehavior int

const (
	// FirstLowerNone does not force the first character to be lowercased.
	FirstLowerNone FirstLowerBehavior = iota
	// FirstLowerAlways forces the first character to be lowercased unconditionally.
	FirstLowerAlways
	// FirstLowerSkipEmpty forces the first character to be lowercased, UNLESS the first parsed word is empty.
	FirstLowerSkipEmpty
)

type caseConfig struct {
	ignore         string
	caseMode       CaseMode
	delimiter      string
	upperIndicator string
	allUpper       bool
	allLower       bool
	screaming      bool
	whispering     bool
	mixCaseSupport bool
	firstUpper     bool
	firstLower     FirstLowerBehavior
	utf8Mode       UTF8Mode
	lowercaseWords map[string]bool
	smartTitleThreshold func(int) float64
}

// OptionIgnore sets characters to be preserved and not considered word boundaries or converted.
func OptionIgnore(ignore string) Option {
	return func(cfg *caseConfig) { cfg.ignore = ignore }
}

// OptionDelimiter sets the delimiter between words.
func OptionDelimiter(d string) Option {
	return func(cfg *caseConfig) { cfg.delimiter = d }
}

// OptionCaseMode sets the case mode.
func OptionCaseMode(caseMode CaseMode) Option {
	return func(cfg *caseConfig) { cfg.caseMode = caseMode }
}

// OptionWhispering ensures all characters are lowercase.
func OptionWhispering() Option {
	return func(cfg *caseConfig) { cfg.caseMode = CMWhispering }
}

// OptionScreaming ensures all characters are uppercase.
func OptionScreaming() Option {
	return func(cfg *caseConfig) { cfg.caseMode = CMScreaming }
}

// OptionFirstUpper ensures the very first character of the result is uppercase.
func OptionFirstUpper() Option {
	return func(cfg *caseConfig) { cfg.firstUpper = true }
}

// OptionFirstLower ensures the very first character of the result is lowercase.
func OptionFirstLower() Option {
	return func(cfg *caseConfig) { cfg.firstLower = FirstLowerAlways }
}

// OptionFirstLowerSkipEmpty behaves like OptionFirstLower, but skips lowercasing if the first parsed word is empty.
func OptionFirstLowerSkipEmpty() Option {
	return func(cfg *caseConfig) { cfg.firstLower = FirstLowerSkipEmpty }
}

// OptionLowercaseWords sets the words to keep lowercase during CMSmartTitle conversion.
func OptionLowercaseWords(words ...string) Option {
	return func(cfg *caseConfig) {
		if cfg.lowercaseWords == nil {
			cfg.lowercaseWords = make(map[string]bool)
		}
		for _, w := range words {
			cfg.lowercaseWords[strings.ToLower(w)] = true
		}
	}
}

// OptionSmartTitleThreshold sets a function that defines the ratio of acronyms to words threshold for fallback to title case.
// For example, if the calculated ratio of acronyms is greater than the threshold returned by this function,
// words will be treated as standard words (e.g. A_NEW_HOPE -> A New Hope) instead of preserving acronym caps.
func OptionSmartTitleThreshold(f func(wordCount int) float64) Option {
	return func(cfg *caseConfig) {
		cfg.smartTitleThreshold = f
	}
}

// OptionMixCaseSupport enables splitting of mixed case words (e.g. CamelCase) into separate words based on uppercase letters.
func OptionMixCaseSupport() Option {
	return func(cfg *caseConfig) { cfg.mixCaseSupport = true }
}

// OptionUpperIndicator sets a specific indicator for upper case (often used for double delimiters).
func OptionUpperIndicator(d string) Option {
	return func(cfg *caseConfig) { cfg.upperIndicator = d }
}

// OptionStrict sets strict mode, which returns an error if invalid UTF-8 sequences are encountered.
func OptionStrict() Option {
	return func(cfg *caseConfig) { cfg.utf8Mode = UTF8Strict }
}

// OptionLoose sets loose mode, which preserves invalid UTF-8 bytes as-is instead of replacing them.
func OptionLoose() Option {
	return func(cfg *caseConfig) { cfg.utf8Mode = UTF8Ignore }
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

	var wordMappers []WordMapper

	for _, opt := range opts {
		switch o := opt.(type) {
		case Option:
			o(cfg)
		case WordMapper:
			wordMappers = append(wordMappers, o)
		}
	}

	for _, m := range wordMappers {
		if m != nil {
			words = m(words)
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
		cfg.firstLower = FirstLowerAlways
	case CMFirstTitle:
		cfg.firstUpper = true
	}


	size := 0
	for _, word := range words {
		if l, ok := word.(interface{ Len() int }); ok {
			size += l.Len() + 5 // +5 heuristic for splitMixCase
		} else {
			size += 10 // fallback
		}
	}
	if len(words) > 1 {
		size += len(cfg.delimiter) * (len(words) - 1)
	}

	var b strings.Builder
	b.Grow(size)

	// Pre-compute first and last non-separator indices for Smart Title mode
	firstNonSep, lastNonSep := -1, -1
	acronymCount := 0
	wordCount := 0
	if cfg.caseMode == CMSmartTitle {
		for i, w := range words {
			if _, ok := w.(SeparatorWord); !ok {
				if firstNonSep == -1 {
					firstNonSep = i
				}
				lastNonSep = i
				wordCount++
				if _, isAcronym := w.(AcronymWord); isAcronym {
					acronymCount++
				} else if _, isUpper := w.(UpperCaseWord); isUpper {
					acronymCount++
				}
			}
		}
	}

	treatAcronymsAsWords := false
	if cfg.smartTitleThreshold != nil && wordCount > 0 {
		if float64(acronymCount)/float64(wordCount) > cfg.smartTitleThreshold(wordCount) {
			treatAcronymsAsWords = true
		}
	}

	var prevIsIgnore bool
	for i, word := range words {
		var isIgnore bool
		if cfg.ignore != "" {
			if s, ok := word.(interface{ String() string }); ok {
				str := s.String()
				if strings.ContainsAny(str, cfg.ignore) {
					isIgnore = true
				}
			}
		}

		if i > 0 && !isIgnore && !prevIsIgnore {
			b.WriteString(cfg.delimiter)
		}
		prevIsIgnore = isIgnore

		switch word := word.(type) {
		case LowercaseWord:
			s := string(word)
			if cfg.allUpper || cfg.screaming {
				for _, r := range s {
					b.WriteRune(unicode.ToUpper(r))
				}
			} else if cfg.caseMode == CMAllTitle {
				var err error
				w, err := upperCaseFirstLower(s, cfg.utf8Mode)
				if err != nil {
					return "", err
				}
				b.WriteString(w)
			} else if cfg.caseMode == CMSmartTitle {
				// Always lower unless first/last
				if i == firstNonSep || i == lastNonSep {
					var err error
					w, err := upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
					b.WriteString(w)
				} else {
					for _, r := range s {
						b.WriteRune(unicode.ToLower(r))
					}
				}
			} else {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			}
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
				var err error
				w, err := upperCaseFirstLower(s, cfg.utf8Mode)
				if err != nil {
					return "", err
				}
				b.WriteString(w)
			} else if cfg.caseMode == CMSmartTitle {
				lowerS := strings.ToLower(s)
				if cfg.lowercaseWords[lowerS] && i != firstNonSep && i != lastNonSep {
					for _, r := range s {
						b.WriteRune(unicode.ToLower(r))
					}
				} else {
					var err error
					w, err := upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
					b.WriteString(w)
				}
			} else {
				for _, r := range s {
					b.WriteRune(unicode.ToLower(r))
				}
			}
		case ExactCaseWord:
			s := string(word)
			if cfg.mixCaseSupport {
				casedDelimiter := cfg.delimiter
				if cfg.allUpper || cfg.screaming {
					casedDelimiter = strings.ToUpper(casedDelimiter)
				} else if cfg.allLower || cfg.whispering {
					casedDelimiter = strings.ToLower(casedDelimiter)
				}
				for j, r := range s {
					if j > 0 && unicode.IsUpper(r) {
						b.WriteString(casedDelimiter)
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
				} else if cfg.caseMode == CMSmartTitle {
					lowerS := strings.ToLower(s)
					if cfg.lowercaseWords[lowerS] && i != firstNonSep && i != lastNonSep {
						for _, r := range s {
							b.WriteRune(unicode.ToLower(r))
						}
					} else {
						var err error
						w, err := upperCaseFirstLower(s, cfg.utf8Mode)
						if err != nil {
							return "", err
						}
						b.WriteString(w)
					}
				} else if cfg.caseMode == CMAllTitle {
					var err error
					w, err := upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
					b.WriteString(w)
				} else {
					b.WriteString(s)
				}
			}
		case FirstUpperCaseWord:
			s := string(word)
			if cfg.caseMode == CMSmartTitle {
				lowerS := strings.ToLower(s)
				if cfg.lowercaseWords[lowerS] && i != firstNonSep && i != lastNonSep {
					s = lowerS
				} else {
					var err error
					s, err = upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
				}
			} else {
				var err error
				s, err = upperCaseFirstLower(s, cfg.utf8Mode)
				if err != nil {
					return "", err
				}
			}
			if cfg.mixCaseSupport {
				casedDelimiter := cfg.delimiter
				if cfg.allUpper || cfg.screaming {
					casedDelimiter = strings.ToUpper(casedDelimiter)
				} else if cfg.allLower || cfg.whispering {
					casedDelimiter = strings.ToLower(casedDelimiter)
				}
				for j, r := range s {
					if j > 0 && unicode.IsUpper(r) {
						b.WriteString(casedDelimiter)
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
				} else if cfg.caseMode == CMSmartTitle {
					lowerS := strings.ToLower(s)
					if cfg.lowercaseWords[lowerS] && i != firstNonSep && i != lastNonSep {
						for _, r := range s {
							b.WriteRune(unicode.ToLower(r))
						}
					} else {
						var err error
						w, err := upperCaseFirstLower(s, cfg.utf8Mode)
						if err != nil {
							return "", err
						}
						b.WriteString(w)
					}
				} else if cfg.caseMode == CMAllTitle {
					var err error
					w, err := upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
					b.WriteString(w)
				} else {
					b.WriteString(s)
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
				var err error
				w, err := upperCaseFirstLower(s, cfg.utf8Mode)
				if err != nil {
					return "", err
				}
				b.WriteString(w)
			} else if cfg.caseMode == CMSmartTitle {
				lowerS := strings.ToLower(s)
				if cfg.lowercaseWords[lowerS] && i != firstNonSep && i != lastNonSep {
					for _, r := range s {
						b.WriteRune(unicode.ToLower(r))
					}
				} else if treatAcronymsAsWords {
					var err error
					w, err := upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
					b.WriteString(w)
				} else {
					b.WriteString(s)
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
				var err error
				w, err := upperCaseFirstLower(s, cfg.utf8Mode)
				if err != nil {
					return "", err
				}
				b.WriteString(w)
			} else if cfg.caseMode == CMSmartTitle {
				lowerS := strings.ToLower(s)
				if cfg.lowercaseWords[lowerS] && i != firstNonSep && i != lastNonSep {
					for _, r := range s {
						b.WriteRune(unicode.ToLower(r))
					}
				} else {
					var err error
					w, err := upperCaseFirstLower(s, cfg.utf8Mode)
					if err != nil {
						return "", err
					}
					b.WriteString(w)
				}
			} else {
				b.WriteString(s)
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
	if cfg.firstLower != FirstLowerNone {
		skipFirstLower := cfg.firstLower == FirstLowerSkipEmpty && len(words) > 0 && words[0].String() == ""
		if !skipFirstLower {
			final = LowerCaseFirst(final)
		}
	}

	return final, nil
}

// PartsToFormattedCase converts Parts to words then formats them.
// This is useful when you have intermediate Parts and want to format them directly.
func PartsToFormattedCase(parts []Part, opts ...any) (string, error) {
	// Extract ParserConfig from opts to use for classification
	p := &ParserConfig{
		SmartAcronyms: true,
		NumberMode:    NumberModeNone,
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
		case ParserOption, Partitioner, PartitionerConfig, SubPartMapper, PartMapper, WordMapper:
			parseOpts = append(parseOpts, v)
		default:
		}
	}
	return parseOpts, fmtOpts
}


// ToKebabCase converts words into kebab-case format.
func ToKebabCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("-")}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToSnakeCase converts words into snake_case format.
func ToSnakeCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("_")}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToPascalCase converts words into PascalCase format.
func ToPascalCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(""), OptionFirstUpper(), OptionCaseMode(CMAllTitle)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToCamelCase converts words into camelCase format.
func ToCamelCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(""), OptionFirstLower(), OptionCaseMode(CMAllTitle)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToDarwinCase converts words into Darwin_Case format (Title_Case with underscore).

// ToLowerCamelCase converts words into lowerCamelCase format.
func ToLowerCamelCase(words []Word, opts ...Option) (string, error) {
	return ToCamelCase(words, opts...)
}

// ToScreamingSnakeCase converts words into SCREAMING_SNAKE_CASE format.
func ToScreamingSnakeCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("_"), OptionCaseMode(CMScreaming)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToScreamingKebabCase converts words into SCREAMING-KEBAB-CASE format.
func ToScreamingKebabCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("-"), OptionCaseMode(CMScreaming)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToDelimitedCase converts words into a string separated by a specific delimiter.
func ToDelimitedCase(words []Word, delimiter uint8, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(string([]byte{delimiter})), OptionCaseMode(CMWhispering)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToScreamingDelimitedCase converts words into a SCREAMING string separated by a specific delimiter.
func ToScreamingDelimitedCase(words []Word, delimiter uint8, ignore string, screaming bool, opts ...Option) (string, error) {
	var mode CaseMode
	if screaming {
		mode = CMScreaming
	} else {
        mode = CMWhispering
    }
	defaults := []any{OptionDelimiter(string([]byte{delimiter})), OptionCaseMode(mode)}
	if ignore != "" {
		defaults = append(defaults, OptionIgnore(ignore))
	}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func ToDarwinCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("_"), OptionCaseMode(CMAllTitle)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// ToTitleCase converts words into a Title Case string (Smart Title).
func ToTitleCase(words []Word, opts ...Option) (string, error) {
	defaults := []any{
		OptionDelimiter(" "),
		OptionFirstUpper(),
		OptionCaseMode(CMSmartTitle),
		OptionLowercaseWords("a", "an", "and", "as", "at", "but", "by", "for", "in", "nor", "of", "on", "or", "so", "the", "to", "yet", "with", "from"),
		OptionSmartTitleThreshold(func(wc int) float64 { return 0.5 }),
	}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// LowercaseWord is a word that will be specifically treated as a minor lowercase word by formatters (e.g., smart title case).
type LowercaseWord string

func (w LowercaseWord) String() string { return string(w) }
func (w LowercaseWord) Len() int       { return len(w) }

// OptionSmartTitleSkipWords is a deprecated alias for OptionLowercaseWords to maintain backward compatibility.
//
// Deprecated: Use OptionLowercaseWords instead.
func OptionSmartTitleSkipWords(words ...string) Option {
	return OptionLowercaseWords(words...)
}
