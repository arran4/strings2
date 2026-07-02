package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

func ExampleWithNumberMode() {
	input := "user123profile"
	// Using WithNumberMode to split at digits
	words, _ := strings2.Parse(input, strings2.WithNumberMode(strings2.NumberModeSplitAlways))
	fmt.Println(words)
	// Output: [user 123 profile]
}

func ExampleWithSmartAcronyms() {
	input := "XMLReader"
	// Using WithSmartAcronyms to separate XML from Reader
	words, _ := strings2.Parse(input, strings2.WithSmartAcronyms(true))
	fmt.Println(words)
	// Output: [XML Reader]
}

func ExampleParseSnakeCase() {
	input := "hello_world"
	// ParseSnakeCase assumes a snake case structure
	words, _ := strings2.ParseSnakeCase(input)
	fmt.Println(words)
	// Output: [hello world]
}
