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
```

### Case Conversion Functions

```go
strings2.ToCamelCase(words)  // "helloWorld"
strings2.ToPascalCase(words) // "HelloWorld"
strings2.ToKebabCase(words)  // "hello-world"
strings2.ToSnakeCase(words)  // "hello_world"
```

### Mapping and Transformation

The `Map` function provides a functional way to process, filter, or transform the sliced `Word`s before formatting.
Mapping can apply arbitrary functions of the type `func([]Word) []Word` to a sequence of words.

```go
import "github.com/arran4/strings2/mappers"

words, _ := strings2.Parse("National Aeronautics and Space Administration")

// Convert words to an acronym
acronymWords := strings2.Map(words, mappers.Acronym)
// Result: [AcronymWord("NAS")]

// Filter elements
noDigits := strings2.Map(words, mappers.Filter(func(w strings2.Word) bool {
    return !strings.ContainsAny(w.String(), "0123456789")
}))

// Reversing words
reversed := strings2.Map(words, mappers.Reverse)
```

Multiple mapping functions can be passed simultaneously and will be applied sequentially.

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
- `OptionFirstUpper()` – force the result to start with an uppercase letter.
- `OptionFirstLower()` – force the result to start with a lowercase letter.

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

strings2 snake --screaming "hello world"
# Result: HELLO_WORLD

strings2 kebab --first-upper "hello world"
# Result: Hello-world
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
