package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

func ExampleToCamelCase() {
	words := []strings2.Word{
		strings2.SingleCaseWord("hello"),
		strings2.SingleCaseWord("world"),
	}
	res, _ := strings2.ToCamelCase(words)
	fmt.Println(res)
	// Output: helloWorld
}

func ExampleToPascalCase() {
	words := []strings2.Word{
		strings2.SingleCaseWord("hello"),
		strings2.SingleCaseWord("world"),
	}
	res, _ := strings2.ToPascalCase(words)
	fmt.Println(res)
	// Output: HelloWorld
}

func ExampleToKebabCase() {
	words := []strings2.Word{
		strings2.SingleCaseWord("hello"),
		strings2.SingleCaseWord("world"),
	}
	res, _ := strings2.ToKebabCase(words)
	fmt.Println(res)
	// Output: hello-world
}

func ExampleToSnakeCase() {
	words := []strings2.Word{
		strings2.SingleCaseWord("hello"),
		strings2.SingleCaseWord("world"),
	}
	res, _ := strings2.ToSnakeCase(words)
	fmt.Println(res)
	// Output: hello_world
}

func ExampleToFormattedCase() {
	words := []strings2.Word{
		strings2.SingleCaseWord("hello"),
		strings2.SingleCaseWord("world"),
	}
	// Screaming snake case
	fmt.Println(strings2.ToFormattedCase(words, strings2.OptionCaseMode(strings2.CMScreaming), strings2.OptionDelimiter("_")))
	// Output: HELLO_WORLD
}
