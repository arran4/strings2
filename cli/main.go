package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"bufio"
	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
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

func buildOpts(delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool) []any {
	var opts []any
	if delimiter != "" {
		opts = append(opts, strings2.OptionDelimiter(delimiter))
		if len(delimiter) > 0 {
			opts = append(opts, strings2.NewPartitioner(strings2.PartitionerConfig{
				Delimiters: map[rune]bool{rune(delimiter[0]): true},
				SplitCamel: true,
			}))
		}
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
	var allAcronyms []string
	if len(acronym) > 0 {
		allAcronyms = append(allAcronyms, acronym...)
	}
	if len(acronymFromFile) > 0 {
		for _, f := range acronymFromFile {
			file, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading acronym file %s: %v\n", f, err)
				os.Exit(1)
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				l := strings.TrimSpace(scanner.Text())
				if l != "" {
					allAcronyms = append(allAcronyms, l)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading acronym file %s: %v\n", f, err)
				file.Close()
				os.Exit(1)
			}
			file.Close()
		}
	}
	if len(allAcronyms) > 0 {
		opts = append(opts, strings2.WordMapper(mappers.MapAcronyms(allAcronyms...)))
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Camel(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Snake(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Kebab(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Pascal(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Darwin(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	process(input, output, args, strings2.ToDarwin, opts...)
}

// Title is a subcommand `strings2 title`
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Title(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	process(input, output, args, strings2.ToTitle, opts...)
}

// LowerCamel is a subcommand `strings2 lowercamel`
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func LowerCamel(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	process(input, output, args, strings2.ToLowerCamel, opts...)
}

// ScreamingSnake is a subcommand `strings2 screamingsnake`
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func ScreamingSnake(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	process(input, output, args, strings2.ToScreamingSnake, opts...)
}

// ScreamingKebab is a subcommand `strings2 screamingkebab`
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
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func ScreamingKebab(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	process(input, output, args, strings2.ToScreamingKebab, opts...)
}

// Delimited is a subcommand `strings2 delimited`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: ".") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Delimited(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	del := uint8('.')
	if len(delimiter) > 0 {
		del = delimiter[0]
	}
	fn := func(in string, options ...any) (string, error) {
		return strings2.ToDelimited(in, del, options...)
	}
	process(input, output, args, fn, opts...)
}

// ScreamingDelimited is a subcommand `strings2 screamingdelimited`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	delimiter: -d --delimiter (default: ".") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	mixCaseSupport: -m --mix-case-support (default: false) Mix case support
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func ScreamingDelimited(input string, output string, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, mixCaseSupport bool, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, mixCaseSupport, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)
	del := uint8('.')
	if len(delimiter) > 0 {
		del = delimiter[0]
	}
	fn := func(in string, options ...any) (string, error) {
		return strings2.ToScreamingDelimited(in, del, "", true, options...)
	}
	process(input, output, args, fn, opts...)
}

func writeJSON(out io.Writer, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(out, string(b))
}

func writeLines[T any](out io.Writer, slice []T, stringer func(T) string) {
	for _, v := range slice {
		fmt.Fprintln(out, stringer(v))
	}
}



type wordData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func getWord(typ string, val string) strings2.Word {
	switch typ {
	case "SingleCaseWord":
		return strings2.SingleCaseWord(val)
	case "FirstUpperCaseWord":
		return strings2.FirstUpperCaseWord(val)
	case "ExactCaseWord":
		return strings2.ExactCaseWord(val)
	case "AcronymWord":
		return strings2.AcronymWord(val)
	case "UpperCaseWord":
		return strings2.UpperCaseWord(val)
	case "SeparatorWord":
		return strings2.SeparatorWord(val)
	default:
		return strings2.ExactCaseWord(val)
	}
}


func writeWords(out io.Writer, words []strings2.Word, jsonOut bool) {
	if jsonOut {
		var wd []wordData
		for _, w := range words {
			wd = append(wd, wordData{Type: fmt.Sprintf("%T", w)[9:], Value: w.String()})
		}
		writeJSON(out, wd)
	} else {
		for _, w := range words {
			fmt.Fprintf(out, "%s:%s\n", fmt.Sprintf("%T", w)[9:], w.String())
		}
	}
}

func getIO(input, output string, args []string) (io.Reader, io.Writer) {
	var in io.Reader
	if input == "-" {
		in = os.Stdin
	} else if input != "" {
		f, err := os.Open(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		in = f
	} else if len(args) > 0 {
		in = strings.NewReader(strings.Join(args, " "))
	} else {
		in = os.Stdin
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
		out = f
	}
	return in, out
}

// Words is a subcommand `strings2 words`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonOut: --json (default: false) Output as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	noSmartAcronyms: --no-smart-acronyms (default: false) Disable smart acronyms
//	acronym: --acronym (default: []) Acronym to preserve case
//	acronymFromFile: --acronym-from-file (default: []) File containing acronyms to preserve case
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Words(input string, output string, jsonOut bool, delimiter string, noSmartAcronyms bool, numberSplitting bool, acronym []string, acronymFromFile []string, strict bool, args ...string) {
	opts := buildOpts(delimiter, false, false, false, false, false, noSmartAcronyms, numberSplitting, acronym, acronymFromFile, strict)

	in, out := getIO(input, output, args)
	b, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if strict && !utf8.ValidString(string(b)) {
		fmt.Fprintf(os.Stderr, "Error: input is not valid UTF-8\n")
		os.Exit(1)
	}

	words, err := strings2.Parse(string(b), opts...)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing: %v\n", err)
		os.Exit(1)
	}

	writeWords(out, words, jsonOut)
}

// Parts is a subcommand `strings2 parts`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonOut: --json (default: false) Output as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	numberSplitting: --number-splitting (default: false) Enable number splitting
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func Parts(input string, output string, jsonOut bool, delimiter string, numberSplitting bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, false, false, false, false, false, false, numberSplitting, nil, nil, strict)

	in, out := getIO(input, output, args)
	b, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if strict && !utf8.ValidString(string(b)) {
		fmt.Fprintf(os.Stderr, "Error: input is not valid UTF-8\n")
		os.Exit(1)
	}

	parts, err := strings2.ParseToParts(string(b), opts...)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing: %v\n", err)
		os.Exit(1)
	}

	if jsonOut {
		writeJSON(out, parts)
	} else {
		writeLines(out, parts, func(w strings2.Part) string { return w.String() })
	}
}

// SubParts is a subcommand `strings2 subparts`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonOut: --json (default: false) Output as JSON
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func SubParts(input string, output string, jsonOut bool, strict bool, args ...string) {
	in, out := getIO(input, output, args)
	b, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	subparts, _ := strings2.StringToSubParts(string(b))

	if jsonOut {
		writeJSON(out, subparts)
	} else {
		writeLines(out, subparts, func(w strings2.SubPart) string { return string(w.Rune()) })
	}
}

func readWords(in io.Reader, jsonInput bool) ([]strings2.Word, error) {
	b, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	var stringWords []string
	if jsonInput {
		if err := json.Unmarshal(b, &stringWords); err != nil {
			return nil, err
		}
	} else {
		// one per line
		lines := strings.Split(string(b), "\n")
		for _, l := range lines {
			if l != "" {
				stringWords = append(stringWords, l)
			}
		}
	}
	var words []strings2.Word
	for _, sw := range stringWords {
		words = append(words, strings2.ExactCaseWord(sw)) // best guess for string->word
	}
	return words, nil
}

func processWordsTo(input string, output string, args []string, jsonInput bool, fn func([]strings2.Word, ...strings2.Option) (string, error), opts ...any) {
	in, out := getIO(input, output, args)
	words, err := readWords(in, jsonInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading words: %v\n", err)
		os.Exit(1)
	}

	// extract strings2.Option
	var sopts []strings2.Option
	for _, o := range opts {
		if so, ok := o.(strings2.Option); ok {
			sopts = append(sopts, so)
		}
	}

	res, err := fn(words, sopts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(out, res)
}

// WordsToCamel is a subcommand `strings2 wordstocamel`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonInput: --json-input (default: false) Input as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func WordsToCamel(input string, output string, jsonInput bool, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, false, false, false, nil, nil, strict)
	processWordsTo(input, output, args, jsonInput, strings2.FromWordsToCamelCase, opts...)
}

// WordsToSnake is a subcommand `strings2 wordstosnake`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonInput: --json-input (default: false) Input as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func WordsToSnake(input string, output string, jsonInput bool, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, false, false, false, nil, nil, strict)
	processWordsTo(input, output, args, jsonInput, strings2.FromWordsToSnakeCase, opts...)
}

// WordsToKebab is a subcommand `strings2 wordstokebab`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonInput: --json-input (default: false) Input as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func WordsToKebab(input string, output string, jsonInput bool, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, false, false, false, nil, nil, strict)
	processWordsTo(input, output, args, jsonInput, strings2.FromWordsToKebabCase, opts...)
}

// WordsToPascal is a subcommand `strings2 wordstopascal`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonInput: --json-input (default: false) Input as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func WordsToPascal(input string, output string, jsonInput bool, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, false, false, false, nil, nil, strict)
	processWordsTo(input, output, args, jsonInput, strings2.FromWordsToPascalCase, opts...)
}

// WordsToDarwin is a subcommand `strings2 wordstodarwin`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonInput: --json-input (default: false) Input as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func WordsToDarwin(input string, output string, jsonInput bool, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, false, false, false, nil, nil, strict)
	processWordsTo(input, output, args, jsonInput, strings2.FromWordsToDarwinCase, opts...)
}

// WordsToTitle is a subcommand `strings2 wordstotitle`
//
// Flags:
//
//	input: -i --input (default: "") Input file or - for stdin
//	output: -o --output (default: "") Output file or - for stdout
//	jsonInput: --json-input (default: false) Input as JSON
//	delimiter: -d --delimiter (default: "") Delimiter
//	screaming: -S --screaming (default: false) Screaming mode
//	whispering: -w --whispering (default: false) Whispering mode
//	firstUpper: -U --first-upper (default: false) First char upper
//	firstLower: -l --first-lower (default: false) First char lower
//	strict: --strict (default: false) Strict UTF8 mode
//	args: ... String to convert if file/stdin not provided
func WordsToTitle(input string, output string, jsonInput bool, delimiter string, screaming bool, whispering bool, firstUpper bool, firstLower bool, strict bool, args ...string) {
	opts := buildOpts(delimiter, screaming, whispering, firstUpper, firstLower, false, false, false, nil, nil, strict)
	processWordsTo(input, output, args, jsonInput, strings2.FromWordsToTitleCase, opts...)
}
