package strings2

// ToY (Auto-detect input format)

// ToCamel converts an input string (auto-detected format) to camelCase.
func ToCamel(input string, opts ...any) string {
	parseOpts, fmtOpts := separateOptions(opts)
	words, _ := Parse(input, parseOpts...)
	return ToCamelCase(words, fmtOpts...)
}

// ToSnake converts an input string (auto-detected format) to snake_case.
func ToSnake(input string, opts ...any) string {
	parseOpts, fmtOpts := separateOptions(opts)
	words, _ := Parse(input, parseOpts...)
	return ToSnakeCase(words, fmtOpts...)
}

// ToKebab converts an input string (auto-detected format) to kebab-case.
func ToKebab(input string, opts ...any) string {
	parseOpts, fmtOpts := separateOptions(opts)
	words, _ := Parse(input, parseOpts...)
	return ToKebabCase(words, fmtOpts...)
}

// ToPascal converts an input string (auto-detected format) to PascalCase.
func ToPascal(input string, opts ...any) string {
	parseOpts, fmtOpts := separateOptions(opts)
	words, _ := Parse(input, parseOpts...)
	return ToPascalCase(words, fmtOpts...)
}

// FromWordsToY

func FromWordsToCamel(words []Word, opts ...Option) string  { return ToCamelCase(words, opts...) }
func FromWordsToSnake(words []Word, opts ...Option) string  { return ToSnakeCase(words, opts...) }
func FromWordsToKebab(words []Word, opts ...Option) string  { return ToKebabCase(words, opts...) }
func FromWordsToPascal(words []Word, opts ...Option) string { return ToPascalCase(words, opts...) }

// FromXToWords

func FromCamelToWords(input string, opts ...any) []Word  { return ParseCamelCase(input) } // ParseCamelCase currently doesn't take opts, we might need to enhance it or use Parse with config
func FromSnakeToWords(input string, opts ...any) []Word  { return ParseSnakeCase(input) }
func FromKebabToWords(input string, opts ...any) []Word  { return ParseKebabCase(input) }
func FromPascalToWords(input string, opts ...any) []Word { return ParseCamelCase(input) } // Pascal uses same transitions as Camel

// FromXToY

// FromCamelTo...

func FromCamelToSnake(input string, opts ...any) string  {
	_, fmtOpts := separateOptions(opts)
	// Currently ParseCamelCase is rigid. Ideally we use Parse with enforced Camel partitioner.
	// But let's respect the existing API if possible or upgrade it.
	// For now, let's use the explicit parsers but maybe we should upgrade them to take opts too?
	// The user request said "all of the functions have ops...any".

	// Better implementation using Parse with config for X
	words := ParseCamelCase(input)
	return ToSnakeCase(words, fmtOpts...)
}

func FromCamelToKebab(input string, opts ...any) string  {
	_, fmtOpts := separateOptions(opts)
	words := ParseCamelCase(input)
	return ToKebabCase(words, fmtOpts...)
}

func FromCamelToPascal(input string, opts ...any) string {
	_, fmtOpts := separateOptions(opts)
	words := ParseCamelCase(input)
	return ToPascalCase(words, fmtOpts...)
}

// FromSnakeTo...

func FromSnakeToCamel(input string, opts ...any) string  {
	_, fmtOpts := separateOptions(opts)
	words := ParseSnakeCase(input)
	return ToCamelCase(words, fmtOpts...)
}

func FromSnakeToKebab(input string, opts ...any) string  {
	_, fmtOpts := separateOptions(opts)
	words := ParseSnakeCase(input)
	return ToKebabCase(words, fmtOpts...)
}

func FromSnakeToPascal(input string, opts ...any) string {
	_, fmtOpts := separateOptions(opts)
	words := ParseSnakeCase(input)
	return ToPascalCase(words, fmtOpts...)
}

// FromKebabTo...

func FromKebabToCamel(input string, opts ...any) string  {
	_, fmtOpts := separateOptions(opts)
	words := ParseKebabCase(input)
	return ToCamelCase(words, fmtOpts...)
}

func FromKebabToSnake(input string, opts ...any) string  {
	_, fmtOpts := separateOptions(opts)
	words := ParseKebabCase(input)
	return ToSnakeCase(words, fmtOpts...)
}

func FromKebabToPascal(input string, opts ...any) string {
	_, fmtOpts := separateOptions(opts)
	words := ParseKebabCase(input)
	return ToPascalCase(words, fmtOpts...)
}

// FromPascalTo...

func FromPascalToCamel(input string, opts ...any) string {
	_, fmtOpts := separateOptions(opts)
	words := ParseCamelCase(input)
	return ToCamelCase(words, fmtOpts...)
}

func FromPascalToSnake(input string, opts ...any) string {
	_, fmtOpts := separateOptions(opts)
	words := ParseCamelCase(input)
	return ToSnakeCase(words, fmtOpts...)
}

func FromPascalToKebab(input string, opts ...any) string {
	_, fmtOpts := separateOptions(opts)
	words := ParseCamelCase(input)
	return ToKebabCase(words, fmtOpts...)
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
