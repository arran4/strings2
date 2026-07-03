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

// StringToWords converts an input string into a slice of Words using the default parser.
// It is an alias for Parse, provided for semantic parity with FromWordsToY functions.
func StringToWords(input string, opts ...any) ([]Word, error) {
	return Parse(input, opts...)
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

// FromCamelToCamel and friends? The user requested permutations.
func FromCamelToCamel(input string, opts ...any) (string, error) {
	words, err := FromCamelToWords(input, opts...)
	if err != nil {
		return "", err
	}
	return FromWordsToCamel(words, extractOptions(opts)...)
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
