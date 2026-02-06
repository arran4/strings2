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

func ExampleNewPartitioner_customFormat() {
	// Inventing a custom format: "Dot.Separated.Values"
	// We want to split by '.' but keep the parts capitalized as is (or handled by classification).
	input := "User.Profile.Settings"

	// Create a partitioner that splits on dot
	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		Delimiters: map[rune]bool{'.': true},
		SplitCamel: true, // Split if there's camel case inside a part
	})

	// Use Parse with the custom partitioner
	words, _ := strings2.Parse(input, partitioner)

	// Convert to Snake Case
	snake := strings2.ToSnakeCase(words)
	fmt.Println(snake)
	// Output: user_profile_settings
}

func ExampleToSnake_options() {
	// Converting Camel to Screaming Snake
	input := "camelCase"
	output, _ := strings2.ToSnake(input, strings2.OptionCaseMode(strings2.CMScreaming))
	fmt.Println(output)
	// Output: CAMEL_CASE
}
