package strings2_test

import (
	"reflect"
	"testing"
	"unicode"

	"github.com/arran4/strings2"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		opts     []strings2.ParserOption
		expected []strings2.Word
	}{
		{
			name:  "CamelCase",
			input: "helloWorld",
			expected: []strings2.Word{
				strings2.SingleCaseWord("hello"),
				strings2.FirstUpperCaseWord("World"),
			},
		},
		{
			name:  "PascalCase",
			input: "HelloWorld",
			expected: []strings2.Word{
				strings2.FirstUpperCaseWord("Hello"),
				strings2.FirstUpperCaseWord("World"),
			},
		},
		{
			name:  "SnakeCase",
			input: "hello_world",
			expected: []strings2.Word{
				strings2.SingleCaseWord("hello"),
				strings2.SingleCaseWord("world"),
			},
		},
		{
			name:  "KebabCase",
			input: "hello-world",
			expected: []strings2.Word{
				strings2.SingleCaseWord("hello"),
				strings2.SingleCaseWord("world"),
			},
		},
		{
			name:  "ScreamingSnakeCase",
			input: "HELLO_WORLD",
			expected: []strings2.Word{
				strings2.AcronymWord("HELLO"),
				strings2.AcronymWord("WORLD"),
			},
		},
		{
			name:  "Sentence",
			input: "This is a brand NEW World",
			expected: []strings2.Word{
				strings2.FirstUpperCaseWord("This"),
				strings2.SingleCaseWord("is"),
				strings2.SingleCaseWord("a"),
				strings2.SingleCaseWord("brand"),
				strings2.AcronymWord("NEW"),
				strings2.FirstUpperCaseWord("World"),
			},
		},
		{
			name:  "Contextual Dash",
			input: "Hello to all the good-doers out there",
			expected: []strings2.Word{
				strings2.FirstUpperCaseWord("Hello"),
				strings2.SingleCaseWord("to"),
				strings2.SingleCaseWord("all"),
				strings2.SingleCaseWord("the"),
				strings2.SingleCaseWord("good-doers"), // Dash not treated as split because Space wins
				strings2.SingleCaseWord("out"),
				strings2.SingleCaseWord("there"),
			},
		},
		{
			name:  "Acronyms with dots",
			input: "N.E.W. World",
			expected: []strings2.Word{
				strings2.AcronymWord("N.E.W."),
				strings2.FirstUpperCaseWord("World"),
			},
		},
		{
			name:  "CamelCase with Acronym",
			input: "PDFLoader",
			expected: []strings2.Word{
				strings2.AcronymWord("PDF"),
				strings2.FirstUpperCaseWord("Loader"),
			},
		},
		{
			name:  "CamelCase with Acronym 2",
			input: "XMLHttp",
			expected: []strings2.Word{
				strings2.AcronymWord("XML"),
				strings2.FirstUpperCaseWord("Http"),
			},
		},
		{
			name:  "Mixed Acronyms",
			input: "jsonXML",
			expected: []strings2.Word{
				strings2.SingleCaseWord("json"),
				strings2.AcronymWord("XML"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strings2.Parse(tt.input, tt.opts...)
			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse() = %#v, want %#v", got, tt.expected)
			}
		})
	}
}

func TestParse_Options(t *testing.T) {
	t.Run("Disable SmartAcronyms", func(t *testing.T) {
		input := "HELLO_WORLD"
		// Default (SmartAcronyms=true) -> AcronymWord
		// Disabled -> UpperCaseWord

		got, _ := strings2.Parse(input, strings2.ParserSmartAcronyms(false))
		expected := []strings2.Word{
			strings2.UpperCaseWord("HELLO"),
			strings2.UpperCaseWord("WORLD"),
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Parse() = %#v, want %#v", got, expected)
		}
	})

	t.Run("Disable SnakeCase Delimiter", func(t *testing.T) {
		input := "hello_world"
		// Remove '_' from delimiters
		got, _ := strings2.Parse(input, strings2.ParserDelimiters{'-', ' ', '.'})

		// '_' is not a delimiter. CamelCase is on by default.
		// "hello_world": all lower. No case change.
		// Result: Single word.

		expected := []strings2.Word{strings2.SingleCaseWord("hello_world")}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Parse() = %#v, want %#v", got, expected)
		}
	})

	t.Run("Custom Delimiter", func(t *testing.T) {
		input := "hello|world"
		got, _ := strings2.Parse(input, strings2.ParserDelimiters{'|'})
		expected := []strings2.Word{
			strings2.SingleCaseWord("hello"),
			strings2.SingleCaseWord("world"),
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Parse() = %#v, want %#v", got, expected)
		}
	})

	t.Run("Custom SubPart Logic", func(t *testing.T) {
		input := "File123Name"

		isDigit := func(r rune) bool { return unicode.IsDigit(r) }

		customSplit := func(pos int, runes []rune) bool {
			if pos == 0 { return false }
			prev := runes[pos-1]
			curr := runes[pos]
			// Split if transitioning from digit to letter or letter to digit
			if isDigit(prev) != isDigit(curr) {
				return true
			}
			// Use default camel case logic too?
			if strings2.DefaultSplitCamelCase(pos, runes) {
				return true
			}
			return false
		}

		got, _ := strings2.Parse(input, strings2.ParserIsNewSubPart(customSplit))

		expected := []strings2.Word{
			strings2.FirstUpperCaseWord("File"),
			strings2.AcronymWord("123"),
			strings2.FirstUpperCaseWord("Name"),
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Parse() = %#v, want %#v", got, expected)
		}
	})
}
