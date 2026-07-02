package main

import (
	"fmt"

	"github.com/arran4/strings2"
)

func main() {
	input := "User.Profile.Settings"

	// Create a partitioner that splits on dot
	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		Delimiters: map[rune]bool{'.': true},
		SplitCamel: true, // Split if there's camel case inside a part
	})

	// Use Parse with the custom partitioner
	words, err := strings2.Parse(input, partitioner)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		return
	}

	fmt.Printf("Parsed words: %v\n", words)

	// Convert to Snake Case
	snake, err := strings2.ToSnakeCase(words, strings2.OptionCaseMode(strings2.CMWhispering))
	if err != nil {
		fmt.Printf("Error formatting: %v\n", err)
		return
	}

	fmt.Printf("Snake Case: %s\n", snake)
}
