package strings2

// ToY (Auto-detect input format)

// ToCamel converts an input string (auto-detected format) to camelCase.
func ToCamel(input string) string {
	words, _ := Parse(input)
	return ToCamelCase(words)
}

// ToSnake converts an input string (auto-detected format) to snake_case.
func ToSnake(input string) string {
	words, _ := Parse(input)
	return ToSnakeCase(words)
}

// ToKebab converts an input string (auto-detected format) to kebab-case.
func ToKebab(input string) string {
	words, _ := Parse(input)
	return ToKebabCase(words)
}

// ToPascal converts an input string (auto-detected format) to PascalCase.
func ToPascal(input string) string {
	words, _ := Parse(input)
	return ToPascalCase(words)
}

// FromWordsToY

func FromWordsToCamel(words []Word) string  { return ToCamelCase(words) }
func FromWordsToSnake(words []Word) string  { return ToSnakeCase(words) }
func FromWordsToKebab(words []Word) string  { return ToKebabCase(words) }
func FromWordsToPascal(words []Word) string { return ToPascalCase(words) }

// FromXToWords

func FromCamelToWords(input string) []Word  { return ParseCamelCase(input) }
func FromSnakeToWords(input string) []Word  { return ParseSnakeCase(input) }
func FromKebabToWords(input string) []Word  { return ParseKebabCase(input) }
func FromPascalToWords(input string) []Word { return ParseCamelCase(input) } // Pascal uses same transitions as Camel

// FromXToY

// FromCamelTo...

func FromCamelToSnake(input string) string  { return ToSnakeCase(ParseCamelCase(input)) }
func FromCamelToKebab(input string) string  { return ToKebabCase(ParseCamelCase(input)) }
func FromCamelToPascal(input string) string { return ToPascalCase(ParseCamelCase(input)) }

// FromSnakeTo...

func FromSnakeToCamel(input string) string  { return ToCamelCase(ParseSnakeCase(input)) }
func FromSnakeToKebab(input string) string  { return ToKebabCase(ParseSnakeCase(input)) }
func FromSnakeToPascal(input string) string { return ToPascalCase(ParseSnakeCase(input)) }

// FromKebabTo...

func FromKebabToCamel(input string) string  { return ToCamelCase(ParseKebabCase(input)) }
func FromKebabToSnake(input string) string  { return ToSnakeCase(ParseKebabCase(input)) }
func FromKebabToPascal(input string) string { return ToPascalCase(ParseKebabCase(input)) }

// FromPascalTo...

func FromPascalToCamel(input string) string { return ToCamelCase(ParseCamelCase(input)) }
func FromPascalToSnake(input string) string { return ToSnakeCase(ParseCamelCase(input)) }
func FromPascalToKebab(input string) string { return ToKebabCase(ParseCamelCase(input)) }
