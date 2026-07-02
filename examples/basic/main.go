package main

import (
	"fmt"

	"github.com/arran4/strings2"
)

func main() {
	input := "hello_world"

	// Convert snake_case to camelCase
	res, err := strings2.FromSnakeToCamel(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Original: %s\n", input)
	fmt.Printf("camelCase: %s\n", res)
}
