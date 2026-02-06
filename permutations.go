package strings2

// ToY (Auto-detect input format)

// ToCamel converts an input string (auto-detected format) to camelCase.
func ToCamel(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := Parse(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToCamelCase(words, fmtOpts...), nil
}

// ToSnake converts an input string (auto-detected format) to snake_case.
func ToSnake(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := Parse(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToSnakeCase(words, fmtOpts...), nil
}

// ToKebab converts an input string (auto-detected format) to kebab-case.
func ToKebab(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := Parse(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToKebabCase(words, fmtOpts...), nil
}

// ToPascal converts an input string (auto-detected format) to PascalCase.
func ToPascal(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := Parse(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToPascalCase(words, fmtOpts...), nil
}

// FromWordsToY

func FromWordsToCamel(words []Word, opts ...Option) string  { return ToCamelCase(words, opts...) }
func FromWordsToSnake(words []Word, opts ...Option) string  { return ToSnakeCase(words, opts...) }
func FromWordsToKebab(words []Word, opts ...Option) string  { return ToKebabCase(words, opts...) }
func FromWordsToPascal(words []Word, opts ...Option) string { return ToPascalCase(words, opts...) }

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

// FromXToY

// FromCamelTo...

func FromCamelToSnake(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseCamelCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToSnakeCase(words, fmtOpts...), nil
}

func FromCamelToKebab(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseCamelCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToKebabCase(words, fmtOpts...), nil
}

func FromCamelToPascal(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseCamelCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToPascalCase(words, fmtOpts...), nil
}

// FromSnakeTo...

func FromSnakeToCamel(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseSnakeCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToCamelCase(words, fmtOpts...), nil
}

func FromSnakeToKebab(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseSnakeCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToKebabCase(words, fmtOpts...), nil
}

func FromSnakeToPascal(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseSnakeCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToPascalCase(words, fmtOpts...), nil
}

// FromKebabTo...

func FromKebabToCamel(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseKebabCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToCamelCase(words, fmtOpts...), nil
}

func FromKebabToSnake(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseKebabCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToSnakeCase(words, fmtOpts...), nil
}

func FromKebabToPascal(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseKebabCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToPascalCase(words, fmtOpts...), nil
}

// FromPascalTo...

func FromPascalToCamel(input string, opts ...any) (string, error) {
	return FromCamelToCamel(input, opts...)
}

func FromPascalToSnake(input string, opts ...any) (string, error) {
	return FromCamelToSnake(input, opts...)
}

func FromPascalToKebab(input string, opts ...any) (string, error) {
	return FromCamelToKebab(input, opts...)
}

func FromCamelToCamel(input string, opts ...any) (string, error) {
	parseOpts, fmtOpts := separateOptions(opts)
	words, err := ParseCamelCase(input, parseOpts...)
	if err != nil {
		return "", err
	}
	return ToCamelCase(words, fmtOpts...), nil
}

// Helper to split ...any into parse opts and format opts
func separateOptions(opts []any) ([]any, []Option) {
	var parseOpts []any
	var fmtOpts []Option

	for _, o := range opts {
		switch v := o.(type) {
		case Option:
			fmtOpts = append(fmtOpts, v)
		case ParserOption, Partitioner, PartitionerConfig:
			parseOpts = append(parseOpts, v)
		default:
			// If it's something else, maybe ignore or panic?
			// For now, assume users pass correct types.
		}
	}
	return parseOpts, fmtOpts
}

// Must swallows error, panics if err != nil
func Must(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}
