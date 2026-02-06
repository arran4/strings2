package strings2_test

import (
	"fmt"
	"strings"

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

func ExampleToFormattedCase_customSpongeCase() {
	// Custom formatting: SpongeBob Case (sPoNgEbOb cAsE)
	// This demonstrates iterating over words and applying custom logic.

	input := "hello world"
	words, _ := strings2.Parse(input)

	var sb strings.Builder
	for i, word := range words {
		if i > 0 {
			sb.WriteString(" ")
		}
		s := word.String()
		for j, r := range s {
			// Simple alternating case logic relative to the whole string start or word start
			// Let's do word-local alternating
			if j%2 == 0 {
				sb.WriteString(strings.ToLower(string(r)))
			} else {
				sb.WriteString(strings.ToUpper(string(r)))
			}
		}
	}

	fmt.Println(sb.String())
	// Output: hElLo wOrLd
}
