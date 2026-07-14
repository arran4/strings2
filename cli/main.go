package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/arran4/strings2"
)

func process(input string, output string, args []string, fn func(string, ...any) (string, error), opts ...any) {
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

	res, err := fn(string(b), opts...)
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

func buildOpts(delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, strict bool) []any {
	var opts []any
	if delimiter != "" {
		opts = append(opts, strings2.OptionDelimiter(delimiter))
	}
	if screaming {
		opts = append(opts, strings2.OptionCaseMode(strings2.CMScreaming))
	}
	if whispering {
		opts = append(opts, strings2.OptionCaseMode(strings2.CMWhispering))
	}
	if firstUpper {
		opts = append(opts, strings2.OptionFirstUpper())
	}
	if firstLower {
		opts = append(opts, strings2.OptionFirstLower())
	}
	if mixCaseSupport {
		opts = append(opts, strings2.OptionMixCaseSupport())
	}
	if noSmartAcronyms {
		opts = append(opts, strings2.WithSmartAcronyms(false))
	}
	if numberSplitting {
		opts = append(opts, strings2.WithNumberSplitting(true))
	}
	if strict {
		opts = append(opts, strings2.OptionStrict())
	}
	return opts
}

// Camel is a subcommand `strings2 camel`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Camel(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, strict)
	process(input, output, args, strings2.ToCamel, opts...)
}

// Snake is a subcommand `strings2 snake`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Snake(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, strict)
	process(input, output, args, strings2.ToSnake, opts...)
}

// Kebab is a subcommand `strings2 kebab`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Kebab(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, strict)
	process(input, output, args, strings2.ToKebab, opts...)
}

// Pascal is a subcommand `strings2 pascal`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Pascal(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, strict)
	process(input, output, args, strings2.ToPascal, opts...)
}

// Darwin is a subcommand `strings2 darwin`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Darwin(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, strict)
	process(input, output, args, strings2.ToDarwin, opts...)
}
