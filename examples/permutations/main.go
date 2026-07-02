package main

import (
	"fmt"

	"github.com/arran4/strings2"
)

func main() {
	input := "mySuperVariable"

	fmt.Printf("Input: %s\n", input)

	// Convert from Camel to Snake
	snake, _ := strings2.FromCamelToSnake(input)
	fmt.Printf("Camel -> Snake: %s\n", snake)

	// Convert from Camel to Kebab
	kebab, _ := strings2.FromCamelToKebab(input)
	fmt.Printf("Camel -> Kebab: %s\n", kebab)

	// Convert from Camel to Pascal
	pascal, _ := strings2.FromCamelToPascal(input)
	fmt.Printf("Camel -> Pascal: %s\n", pascal)

	// Chain conversion back to Snake to demonstrate it handles different casing properly
	snake2, _ := strings2.FromPascalToSnake(pascal)
	fmt.Printf("Pascal -> Snake: %s\n", snake2)
}
