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

func TestExplicitParsers(t *testing.T) {
	t.Run("ParseSnakeCase", func(t *testing.T) {
		got := strings2.ParseSnakeCase("hello_world")
		expected := []strings2.Word{
			strings2.SingleCaseWord("hello"),
			strings2.SingleCaseWord("world"),
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("ParseSnakeCase() = %#v, want %#v", got, expected)
		}
	})

	t.Run("ParseCamelCase", func(t *testing.T) {
		got := strings2.ParseCamelCase("helloWorld")
		expected := []strings2.Word{
			strings2.SingleCaseWord("hello"),
			strings2.FirstUpperCaseWord("World"),
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("ParseCamelCase() = %#v, want %#v", got, expected)
		}
	})
}
