package strings2

// ToY (Auto-detect input format)

// ToCamel converts an input string (auto-detected format) to camelCase.
func ToCamel(input string, opts ...any) (string, error) {
	// Camel: Delimiter "", FirstLower, AllTitle
	defaults := []any{OptionDelimiter(""), OptionFirstLower(), OptionCaseMode(CMAllTitle)}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToSnake converts an input string (auto-detected format) to snake_case.
func ToSnake(input string, opts ...any) (string, error) {
	// Snake: Delimiter "_"
	defaults := []any{OptionDelimiter("_")}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToKebab converts an input string (auto-detected format) to kebab-case.
func ToKebab(input string, opts ...any) (string, error) {
	// Kebab: Delimiter "-"
	defaults := []any{OptionDelimiter("-")}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToPascal converts an input string (auto-detected format) to PascalCase.
func ToPascal(input string, opts ...any) (string, error) {
	// Pascal: Delimiter "", FirstUpper, AllTitle
	defaults := []any{OptionDelimiter(""), OptionFirstUpper(), OptionCaseMode(CMAllTitle)}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToTitle converts an input string (auto-detected format) to Title Case (Smart Title).
func ToTitle(input string, opts ...any) (string, error) {
	// Title: Delimiter " ", FirstUpper, SmartTitle, Default Skip Words
	defaults := []any{OptionDelimiter(" "), OptionFirstUpper(), OptionCaseMode(CMSmartTitle), OptionSmartTitleSkipWords("a", "an", "and", "as", "at", "but", "by", "for", "in", "nor", "of", "on", "or", "so", "the", "to", "yet", "with", "from"), OptionSmartTitleThreshold(func(wc int) float64 { return 0.5 })}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToDarwin converts an input string (auto-detected format) to Darwin_Case.

// ToLowerCamel converts an input string (auto-detected format) to lowerCamelCase.
func ToLowerCamel(input string, opts ...any) (string, error) {
	return ToCamel(input, opts...)
}

// ToScreamingSnake converts an input string (auto-detected format) to SCREAMING_SNAKE_CASE.
func ToScreamingSnake(input string, opts ...any) (string, error) {
	defaults := []any{OptionDelimiter("_"), OptionCaseMode(CMScreaming)}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToScreamingKebab converts an input string (auto-detected format) to SCREAMING-KEBAB-CASE.
func ToScreamingKebab(input string, opts ...any) (string, error) {
	defaults := []any{OptionDelimiter("-"), OptionCaseMode(CMScreaming)}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToDelimited converts an input string (auto-detected format) to a string separated by a specific delimiter.
func ToDelimited(input string, delimiter uint8, opts ...any) (string, error) {
	defaults := []any{OptionDelimiter(string([]byte{delimiter})), OptionCaseMode(CMWhispering)}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// ToScreamingDelimited converts an input string (auto-detected format) to a SCREAMING string separated by a specific delimiter.
func ToScreamingDelimited(input string, delimiter uint8, ignore string, screaming bool, opts ...any) (string, error) {
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
	return ToFormattedString(input, append(defaults, opts...)...)
}

func ToDarwin(input string, opts ...any) (string, error) {
	defaults := []any{OptionDelimiter("_"), OptionCaseMode(CMAllTitle)}
	return ToFormattedString(input, append(defaults, opts...)...)
}

// FromWordsToY

// FromWordsToCamel converts words to camelCase.
func FromWordsToCamel(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(""), OptionFirstLower(), OptionCaseMode(CMAllTitle)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToSnake(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("_")}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToKebab(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("-")}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToPascal(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(""), OptionFirstUpper(), OptionCaseMode(CMAllTitle)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToTitle(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(" "), OptionFirstUpper(), OptionCaseMode(CMSmartTitle), OptionSmartTitleSkipWords("a", "an", "and", "as", "at", "but", "by", "for", "in", "nor", "of", "on", "or", "so", "the", "to", "yet", "with", "from"), OptionSmartTitleThreshold(func(wc int) float64 { return 0.5 })}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}


func FromWordsToLowerCamel(words []Word, opts ...Option) (string, error) {
	return FromWordsToCamel(words, opts...)
}

func FromWordsToScreamingSnake(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("_"), OptionCaseMode(CMScreaming)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToScreamingKebab(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("-"), OptionCaseMode(CMScreaming)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToDelimited(words []Word, delimiter uint8, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter(string([]byte{delimiter})), OptionCaseMode(CMWhispering)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

func FromWordsToScreamingDelimited(words []Word, delimiter uint8, ignore string, screaming bool, opts ...Option) (string, error) {
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

func FromWordsToDarwin(words []Word, opts ...Option) (string, error) {
	defaults := []any{OptionDelimiter("_"), OptionCaseMode(CMAllTitle)}
	return WordsToFormattedCase(words, append(defaults, convertOptions(opts)...)...)
}

// FromXToWords

func FromCamelToWords(input string, opts ...any) ([]Word, error) {
	return ParseCamelCase(input, opts...)
}
func FromSnakeToWords(input string, opts ...any) ([]Word, error) {
	return ParseSnakeCase(input, opts...)
}
func FromKebabToWords(input string, opts ...any) ([]Word, error) {
	return ParseKebabCase(input, opts...)
}
func FromPascalToWords(input string, opts ...any) ([]Word, error) {
	return ParseCamelCase(input, opts...)
}
func FromDarwinToWords(input string, opts ...any) ([]Word, error) {
	return ParseDarwinCase(input, opts...)
}

// FromXToY

// FromCamelTo...

func FromCamelToSnake(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToSnake(words, extractOptions(opts)...)
}

func FromCamelToKebab(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToKebab(words, extractOptions(opts)...)
}

func FromCamelToPascal(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToPascal(words, extractOptions(opts)...)
}

func FromCamelToTitle(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToTitle(words, extractOptions(opts)...)
}

func FromCamelToDarwin(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToDarwin(words, extractOptions(opts)...)
}

// FromSnakeTo...

func FromSnakeToCamel(input string, opts ...any) (string, error) {
	words, err := FromSnakeToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToCamel(words, extractOptions(opts)...)
}

func FromSnakeToKebab(input string, opts ...any) (string, error) {
	words, err := FromSnakeToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToKebab(words, extractOptions(opts)...)
}

func FromSnakeToPascal(input string, opts ...any) (string, error) {
	words, err := FromSnakeToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToPascal(words, extractOptions(opts)...)
}

func FromSnakeToTitle(input string, opts ...any) (string, error) {
	words, err := FromSnakeToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToTitle(words, extractOptions(opts)...)
}

func FromSnakeToDarwin(input string, opts ...any) (string, error) {
	words, err := FromSnakeToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToDarwin(words, extractOptions(opts)...)
}

// FromKebabTo...

func FromKebabToCamel(input string, opts ...any) (string, error) {
	words, err := FromKebabToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToCamel(words, extractOptions(opts)...)
}

func FromKebabToSnake(input string, opts ...any) (string, error) {
	words, err := FromKebabToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToSnake(words, extractOptions(opts)...)
}

func FromKebabToPascal(input string, opts ...any) (string, error) {
	words, err := FromKebabToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToPascal(words, extractOptions(opts)...)
}

func FromKebabToTitle(input string, opts ...any) (string, error) {
	words, err := FromKebabToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToTitle(words, extractOptions(opts)...)
}

func FromKebabToDarwin(input string, opts ...any) (string, error) {
	words, err := FromKebabToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToDarwin(words, extractOptions(opts)...)
}

// FromPascalTo...

func FromPascalToCamel(input string, opts ...any) (string, error) {
	words, err := FromPascalToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToCamel(words, extractOptions(opts)...)
}

func FromPascalToSnake(input string, opts ...any) (string, error) {
	words, err := FromPascalToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToSnake(words, extractOptions(opts)...)
}

func FromPascalToKebab(input string, opts ...any) (string, error) {
	words, err := FromPascalToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToKebab(words, extractOptions(opts)...)
}

func FromPascalToTitle(input string, opts ...any) (string, error) {
	words, err := FromPascalToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToTitle(words, extractOptions(opts)...)
}

func FromPascalToDarwin(input string, opts ...any) (string, error) {
	words, err := FromPascalToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToDarwin(words, extractOptions(opts)...)
}

// FromCamelToCamel and friends? The user requested permutations.
func FromCamelToCamel(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToCamel(words, extractOptions(opts)...)
}

// FromDarwinTo...

func FromDarwinToCamel(input string, opts ...any) (string, error) {
	words, err := FromDarwinToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToCamel(words, extractOptions(opts)...)
}

func FromDarwinToSnake(input string, opts ...any) (string, error) {
	words, err := FromDarwinToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToSnake(words, extractOptions(opts)...)
}

func FromDarwinToKebab(input string, opts ...any) (string, error) {
	words, err := FromDarwinToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToKebab(words, extractOptions(opts)...)
}

func FromDarwinToPascal(input string, opts ...any) (string, error) {
	words, err := FromDarwinToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToPascal(words, extractOptions(opts)...)
}

func FromDarwinToTitle(input string, opts ...any) (string, error) {
	words, err := FromDarwinToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToTitle(words, extractOptions(opts)...)
}

// extractOptions helpers to get just []Option for formatters
func extractOptions(opts []any) []Option {
	var out []Option
	for _, o := range opts {
		if opt, ok := o.(Option); ok {
			out = append(out, opt)
		}
	}
	return out
}

// WordMappers shouldn't need extraction as they are processed in FromXToY natively if they are passed as part of the parse chain, or in WordsToFormattedCase if passed.

// Helper to separate options - mostly used internally if needed, but here we delegate.
// Keeping Must as well.

// Must swallows error, panics if err != nil
func Must(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}
