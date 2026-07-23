# strings2

[![CI Status](https://github.com/arran4/strings2/actions/workflows/ci.yml/badge.svg)](https://github.com/arran4/strings2/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/arran4/strings2.svg)](https://pkg.go.dev/github.com/arran4/strings2)

strings2 provides utilities for converting slices of words into various casing conventions. It is intended to supplement Go's standard library `strings` package with helpers for creating formats such as `camelCase`, `PascalCase`, `snake_case` and `kebab-case`.

## Installation

```
go get github.com/arran4/strings2
```

Add the module to your project and import it:

```go
import "github.com/arran4/strings2"
```

## Usage

Words must implement `fmt.Stringer`. The package defines several helper types which satisfy this interface:

```go
words := []strings2.Word{
    strings2.SingleCaseWord("hello"),
    strings2.SingleCaseWord("world"),
}
```

### Parsing

The library includes a robust parser to convert strings into typed `Word` objects, distinguishing between acronyms, casing, and delimiters.

```go
// Auto-detect format and parse
words, err := strings2.Parse("helloWorld")
// Result: [SingleCaseWord("hello"), FirstUpperCaseWord("World")]

// Parse specific format
words = strings2.ParseSnakeCase("hello_world")

// Configure parser
words, err = strings2.Parse("N.E.W. World", strings2.ParserSmartAcronyms(true))

// Configure parser to ignore characters, preventing them from being consumed as delimiters
words, err = strings2.Parse("foo.bar", strings2.WithIgnore("."))

// Custom multi-byte delimiters using DelimiterDetector
partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
	DelimiterDetector: func(subs []strings2.SubPart, index int) int {
		if index+1 < len(subs) && subs[index].Rune() == ':' && subs[index+1].Rune() == ':' {
			return 2 // matched "::"
		}
		return 0
	},
})
words, err = strings2.Parse("foo::bar", partitioner)
```

### Case Conversion Functions

```go
strings2.ToCamelCase(words)  // "helloWorld"
strings2.ToLowerCamelCase(words) // "helloWorld"
strings2.ToPascalCase(words) // "HelloWorld"
strings2.ToKebabCase(words)  // "hello-world"
strings2.ToScreamingKebabCase(words) // "HELLO-WORLD"
strings2.ToSnakeCase(words)  // "hello_world"
strings2.ToScreamingSnakeCase(words) // "HELLO_WORLD"
strings2.ToDarwinCase(words) // "Hello_World"
titleWords, _ := strings2.Parse("the lord of the rings")
strings2.ToTitleCase(titleWords)  // "The Lord of the Rings"
strings2.ToDelimitedCase(words, '.') // "hello.world"
strings2.ToScreamingDelimitedCase(words, '.', "", true) // "HELLO.WORLD"
strings2.ToScreamingDelimitedCase(words, '_', ".", true) // "HELLO.WORLD" (where "." is ignored and preserved)
```

### Mapping and Transformation

You can provide `WordMapper` (`func([]Word) []Word`), `PartMapper`, or `SubPartMapper` natively into formatting functions or `Parse`/`StringToWords` as variadic options to filter, transform, or reorder elements during parsing and generation.

```go
import "github.com/arran4/strings2/mappers"

// Convert strings natively with inline options
acronym, _ := strings2.ToFormattedString(
    "National Aeronautics and Space Administration",
    strings2.WordMapper(mappers.Acronym),
    strings2.OptionCaseMode(strings2.CMVerbatim),
    strings2.OptionDelimiter(""),
)
// Result: "NAASA"

// Preserve lowercase words using a lowercase predicate
mappedTitle, _ := strings2.ToTitleCase(
    titleWords,
    strings2.WordMapper(mappers.MapLowercase("via")),
)
// Result: "The Lord of the Rings via Middle Earth"

// Reversing words natively
reversed, _ := strings2.ToCamel("hello world from strings2", strings2.WordMapper(mappers.Reverse))
// Result: "strings2FromWorldHello"

// Filter out numbers natively
filterNumbers := mappers.Filter(func(w strings2.Word) bool {
    return !strings.ContainsAny(w.String(), "0123456789")
})
noDigits, _ := strings2.ToSnake("hello 123 world", strings2.WordMapper(filterNumbers))
// Result: "hello_world"
```

Multiple mapping functions can be passed simultaneously and will be applied sequentially in their respective lifecycle phase (SubPart, Part, then Word).

### Customising Formatting

Behaviour can be tuned with options passed to each function. Some commonly used options include:

- `OptionDelimiter(string)` – change the delimiter used between words.
- `OptionCaseMode(CaseMode)` – set the case transformation mode. Modes include:
  - `CMVerbatim`
  - `CMFirstTitle`
  - `CMAllTitle`
  - `CMFirstLower`
  - `CMWhispering`
  - `CMScreaming`
  - `CMSmartTitle`
- `OptionFirstUpper()` – force the result to start with an uppercase letter.
- `OptionFirstLower()` – force the result to start with a lowercase letter.
- `OptionSmartTitleUpperMode(SmartTitleUpperMode)` – force normalization or preservation of uppercase words in `CMSmartTitle`.
- `OptionSmartTitleAcronymPredicate(func(string) bool)` – provide custom knowledge for matching domain acronyms.

Examples:

```go
// Custom delimiter
fmt.Println(strings2.ToKebabCase(words, strings2.OptionDelimiter("|")))

// Screaming snake case
fmt.Println(strings2.ToSnakeCase(words, strings2.OptionCaseMode(strings2.CMScreaming)))
```

### CLI Mode

The library also provides a command-line interface that exposes all these options, ensuring that the CLI mode has as much flexibility as the code (without being obligated to use smart defaults).

```bash
strings2 camel "hello world"
# Result: helloWorld

strings2 lowercamel "hello world"
# Result: helloWorld

strings2 snake --screaming "hello world"
# Result: HELLO_WORLD

strings2 screamingsnake "hello world"
# Result: HELLO_WORLD

strings2 darwin "hello world"
# Result: Hello_World

strings2 kebab --first-upper "hello world"
# Result: Hello-world

strings2 screamingkebab "hello world"
# Result: HELLO-WORLD

strings2 delimited "hello world" -d "."
# Result: hello.world

strings2 screamingdelimited "hello world" -d "."
# Result: HELLO.WORLD
```

You can pipe input into the CLI as well:
```bash
echo "hello world" | strings2 pascal
# Result: HelloWorld
```

Available flags across commands:
- `--delimiter`, `-d` (string): Override the delimiter
- `--screaming`, `-S`: Enforce uppercase formatting
- `--whispering`, `-w`: Enforce lowercase formatting
- `--first-upper`, `-U`: Capitalize the first letter
- `--first-lower`, `-l`: Lowercase the first letter
- `--mix-case-support`, `-m`: Enable splitting of mixed case words
- `--no-smart-acronyms`: Disable acronym preservation
- `--number-splitting`: Enable letter-digit boundary splitting
```

Options are composable so multiple behaviours can be applied at once. See the documentation in `types.go` for details on further options.

## TODO

- Support slices for flags when the gosubc version supports it.

## License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.
