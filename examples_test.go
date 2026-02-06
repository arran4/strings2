package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

// Example usage tests (documentation tests)

func ExampleParse() {
	words, _ := strings2.Parse("helloWorld")
	fmt.Println(words)
	// Output: [hello World]
}

func ExampleParse_smartAcronyms() {
	words, _ := strings2.Parse("XMLReader", strings2.WithSmartAcronyms(true))
	fmt.Println(words)
	// Output: [XML Reader]
}

func ExampleParse_snakeCase() {
	words, _ := strings2.ParseSnakeCase("hello_world")
	fmt.Println(words)
	// Output: [hello world]
}
