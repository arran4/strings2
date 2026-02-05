package strings2_test

import (
	"reflect"
	"testing"

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

	t.Run("Restricted Format", func(t *testing.T) {
		input := "hello_world" // Snake
		// Allow only Camel
		got, _ := strings2.Parse(input, strings2.ParserAllowedFormats(strings2.FormatCamelCase))
		// Should detect Unknown or fail to match Snake logic.
		// Fallback to single string.
		// "hello_world" -> SingleCaseWord (all lower letters)

		expected := []strings2.Word{strings2.SingleCaseWord("hello_world")}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Parse() = %#v, want %#v", got, expected)
		}
	})
}
