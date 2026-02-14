package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/arran4/strings2"
)

func process(input string, output string, args []string, fn func(string, ...any) (string, error)) {
	var in io.Reader
	if input == "-" {
		in = os.Stdin
	} else if input != "" {
		f, err := os.Open(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	} else if len(args) > 0 {
		in = strings.NewReader(strings.Join(args, " "))
	} else {
		in = os.Stdin
	}

	b, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	res, err := fn(string(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing: %v\n", err)
		os.Exit(1)
	}

	var out io.Writer
	if output == "-" || output == "" {
		out = os.Stdout
	} else {
		f, err := os.Create(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		out = f
	}

	fmt.Fprintln(out, res)
}

// Camel is a subcommand `strings2 camel`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	args: ... Positional arguments
func Camel(input string, output string, args ...string) {
	process(input, output, args, strings2.ToCamel)
}

// Snake is a subcommand `strings2 snake`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	args: ... Positional arguments
func Snake(input string, output string, args ...string) {
	process(input, output, args, strings2.ToSnake)
}

// Kebab is a subcommand `strings2 kebab`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	args: ... Positional arguments
func Kebab(input string, output string, args ...string) {
	process(input, output, args, strings2.ToKebab)
}

// Pascal is a subcommand `strings2 pascal`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	args: ... Positional arguments
func Pascal(input string, output string, args ...string) {
	process(input, output, args, strings2.ToPascal)
}
