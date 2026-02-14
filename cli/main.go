package cli

import (
	"fmt"
	"os"

	"github.com/arran4/strings2"
)

// Camel is a subcommand `strings2 camel`
//
// Flags:
//
//	input: @1 Input string
func Camel(input string) {
	res, err := strings2.ToCamel(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}

// Snake is a subcommand `strings2 snake`
//
// Flags:
//
//	input: @1 Input string
func Snake(input string) {
	res, err := strings2.ToSnake(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}

// Kebab is a subcommand `strings2 kebab`
//
// Flags:
//
//	input: @1 Input string
func Kebab(input string) {
	res, err := strings2.ToKebab(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}

// Pascal is a subcommand `strings2 pascal`
//
// Flags:
//
//	input: @1 Input string
func Pascal(input string) {
	res, err := strings2.ToPascal(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}
