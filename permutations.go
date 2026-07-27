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
	defaults := []any{OptionDelimiter(" "), OptionFirstUpper(), OptionCaseMode(CMSmartTitle), OptionLowercaseWords("a", "an", "and", "as", "at", "but", "by", "for", "in", "nor", "of", "on", "or", "so", "the", "to", "yet", "with", "from")}
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
	defaults := []any{OptionDelimiter(" "), OptionFirstUpper(), OptionCaseMode(CMSmartTitle), OptionLowercaseWords("a", "an", "and", "as", "at", "but", "by", "for", "in", "nor", "of", "on", "or", "so", "the", "to", "yet", "with", "from")}
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

// --- Global parsers ---
func ToWords(input string, opts ...any) ([]Word, error)           { return Parse(input, opts...) }
func FromStringToWords(input string, opts ...any) ([]Word, error) { return Parse(input, opts...) }
func ToParts(input string, opts ...any) ([]Part, error)           { return ParseToParts(input, opts...) }
func FromStringToParts(input string, opts ...any) ([]Part, error) {
	return ParseToParts(input, opts...)
}
func ParseToSubParts(input string) ([]SubPart, Stats)      { return StringToSubParts(input) }
func ToSubParts(input string) ([]SubPart, Stats)           { return StringToSubParts(input) }
func FromStringToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }

// --- FromWordsTo[Format]Case ---
func FromWordsToCamelCase(words []Word, opts ...Option) (string, error) {
	return ToCamelCase(words, opts...)
}
func FromWordsToSnakeCase(words []Word, opts ...Option) (string, error) {
	return ToSnakeCase(words, opts...)
}
func FromWordsToKebabCase(words []Word, opts ...Option) (string, error) {
	return ToKebabCase(words, opts...)
}
func FromWordsToPascalCase(words []Word, opts ...Option) (string, error) {
	return ToPascalCase(words, opts...)
}
func FromWordsToDarwinCase(words []Word, opts ...Option) (string, error) {
	return ToDarwinCase(words, opts...)
}
func FromWordsToTitleCase(words []Word, opts ...Option) (string, error) {
	return ToTitleCase(words, opts...)
}
func FromWordsToLowerCamelCase(words []Word, opts ...Option) (string, error) {
	return ToLowerCamelCase(words, opts...)
}
func FromWordsToScreamingSnakeCase(words []Word, opts ...Option) (string, error) {
	return ToScreamingSnakeCase(words, opts...)
}
func FromWordsToScreamingKebabCase(words []Word, opts ...Option) (string, error) {
	return ToScreamingKebabCase(words, opts...)
}
func FromWordsToDelimitedCase(words []Word, delimiter uint8, opts ...Option) (string, error) {
	return ToDelimitedCase(words, delimiter, opts...)
}
func FromWordsToScreamingDelimitedCase(words []Word, delimiter uint8, ignore string, screaming bool, opts ...Option) (string, error) {
	return ToScreamingDelimitedCase(words, delimiter, ignore, screaming, opts...)
}

// --- From[Format]Case[To...] ---
func FromCamelCase(input string, opts ...any) ([]Word, error) { return ParseCamelCase(input, opts...) }
func FromCamelCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseCamelCase(input, opts...)
}
func FromCamelCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseCamelCaseToParts(input, opts...)
}
func FromCamelCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromSnakeCase(input string, opts ...any) ([]Word, error) { return ParseSnakeCase(input, opts...) }
func FromSnakeCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseSnakeCase(input, opts...)
}
func FromSnakeCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseSnakeCaseToParts(input, opts...)
}
func FromSnakeCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromKebabCase(input string, opts ...any) ([]Word, error) { return ParseKebabCase(input, opts...) }
func FromKebabCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseKebabCase(input, opts...)
}
func FromKebabCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseKebabCaseToParts(input, opts...)
}
func FromKebabCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromPascalCase(input string, opts ...any) ([]Word, error) {
	return ParsePascalCase(input, opts...)
}
func FromPascalCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParsePascalCase(input, opts...)
}
func FromPascalCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParsePascalCaseToParts(input, opts...)
}
func FromPascalCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromDarwinCase(input string, opts ...any) ([]Word, error) {
	return ParseDarwinCase(input, opts...)
}
func FromDarwinCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseDarwinCase(input, opts...)
}
func FromDarwinCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseDarwinCaseToParts(input, opts...)
}
func FromDarwinCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromTitleCase(input string, opts ...any) ([]Word, error)  { return ParseTitleCase(input, opts...) }
func FromTitleCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseTitleCase(input, opts...)
}
func FromTitleCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseTitleCaseToParts(input, opts...)
}
func FromTitleCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromLowerCamelCase(input string, opts ...any) ([]Word, error) {
	return ParseLowerCamelCase(input, opts...)
}
func FromLowerCamelCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseLowerCamelCase(input, opts...)
}
func FromLowerCamelCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseLowerCamelCaseToParts(input, opts...)
}
func FromLowerCamelCaseToSubParts(input string) ([]SubPart, Stats) { return StringToSubParts(input) }
func FromScreamingSnakeCase(input string, opts ...any) ([]Word, error) {
	return ParseScreamingSnakeCase(input, opts...)
}
func FromScreamingSnakeCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseScreamingSnakeCase(input, opts...)
}
func FromScreamingSnakeCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseScreamingSnakeCaseToParts(input, opts...)
}
func FromScreamingSnakeCaseToSubParts(input string) ([]SubPart, Stats) {
	return StringToSubParts(input)
}
func FromScreamingKebabCase(input string, opts ...any) ([]Word, error) {
	return ParseScreamingKebabCase(input, opts...)
}
func FromScreamingKebabCaseToWords(input string, opts ...any) ([]Word, error) {
	return ParseScreamingKebabCase(input, opts...)
}
func FromScreamingKebabCaseToParts(input string, opts ...any) ([]Part, error) {
	return ParseScreamingKebabCaseToParts(input, opts...)
}
func FromScreamingKebabCaseToSubParts(input string) ([]SubPart, Stats) {
	return StringToSubParts(input)
}

// --- FromStringTo[Format] ---
func FromStringToCamel(input string, opts ...any) (string, error)     { return ToCamel(input, opts...) }
func FromStringToCamelCase(input string, opts ...any) (string, error) { return ToCamel(input, opts...) }
func FromStringToSnake(input string, opts ...any) (string, error)     { return ToSnake(input, opts...) }
func FromStringToSnakeCase(input string, opts ...any) (string, error) { return ToSnake(input, opts...) }
func FromStringToKebab(input string, opts ...any) (string, error)     { return ToKebab(input, opts...) }
func FromStringToKebabCase(input string, opts ...any) (string, error) { return ToKebab(input, opts...) }
func FromStringToPascal(input string, opts ...any) (string, error)    { return ToPascal(input, opts...) }
func FromStringToPascalCase(input string, opts ...any) (string, error) {
	return ToPascal(input, opts...)
}
func FromStringToDarwin(input string, opts ...any) (string, error) { return ToDarwin(input, opts...) }
func FromStringToDarwinCase(input string, opts ...any) (string, error) {
	return ToDarwin(input, opts...)
}
func FromStringToTitle(input string, opts ...any) (string, error)     { return ToTitle(input, opts...) }
func FromStringToTitleCase(input string, opts ...any) (string, error) { return ToTitle(input, opts...) }
func FromStringToLowerCamel(input string, opts ...any) (string, error) {
	return ToLowerCamel(input, opts...)
}
func FromStringToLowerCamelCase(input string, opts ...any) (string, error) {
	return ToLowerCamel(input, opts...)
}
func FromStringToScreamingSnake(input string, opts ...any) (string, error) {
	return ToScreamingSnake(input, opts...)
}
func FromStringToScreamingSnakeCase(input string, opts ...any) (string, error) {
	return ToScreamingSnake(input, opts...)
}
func FromStringToScreamingKebab(input string, opts ...any) (string, error) {
	return ToScreamingKebab(input, opts...)
}
func FromStringToScreamingKebabCase(input string, opts ...any) (string, error) {
	return ToScreamingKebab(input, opts...)
}
func FromStringToDelimited(input string, delimiter uint8, opts ...any) (string, error) {
	return ToDelimited(input, delimiter, opts...)
}
func FromStringToDelimitedCase(input string, delimiter uint8, opts ...any) (string, error) {
	return ToDelimited(input, delimiter, opts...)
}
func FromStringToScreamingDelimited(input string, delimiter uint8, ignore string, screaming bool, opts ...any) (string, error) {
	return ToScreamingDelimited(input, delimiter, ignore, screaming, opts...)
}
func FromStringToScreamingDelimitedCase(input string, delimiter uint8, ignore string, screaming bool, opts ...any) (string, error) {
	return ToScreamingDelimited(input, delimiter, ignore, screaming, opts...)
}
