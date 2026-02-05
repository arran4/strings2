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
		// "Hello to all the good-doers out there"
		// This has both spaces and hyphens.
		// If "Sentence" logic is picked (stats.Spaces > 0), it splits by space.
		// Result: "Hello", "to", "all", "the", "good-doers", "out", "there".
		// But in the previous failure, the output was:
		// []strings2.Word{"Hello to all the good", "doers out there"}
		// This suggests it picked KebabCasePartitioner!
		// Why?
		// DetectPartitioner:
		// if stats.SymbolCounts['_'] > 0 { return SnakeCasePartitioner }
		// if stats.SymbolCounts['-'] > 0 { return KebabCasePartitioner }
		// if stats.Spaces > 0 { return SpacePartitioner }
		//
		// Order matters! It checks Hyphen BEFORE Space.
		// So if input has hyphens, it picks KebabCase.
		// KebabCase splits on '-'.
		// "Hello to all the good-doers out there" -> "Hello to all the good", "doers out there"
		// This seems wrong for sentences.
		// We should probably prioritize Sentence (Space) over Kebab (Hyphen) if both exist?
		// Or refine the heuristic.
		// But for now, let's just fix the test expectation to match the current logic or remove ambiguous test.
		// Wait, user asked for "Contextual".
		// If I change the order in `DetectPartitioner` to Space first, then KebabCase won't work if it has spaces?
		// Real KebabCase shouldn't have spaces.
		// So if Spaces > 0, it's probably NOT KebabCase (strict).
		// So moving Space check up seems correct.

		{
			name:  "Contextual Dash",
			input: "Hello to all the good-doers out there",
			expected: []strings2.Word{
				strings2.FirstUpperCaseWord("Hello"),
				strings2.SingleCaseWord("to"),
				strings2.SingleCaseWord("all"),
				strings2.SingleCaseWord("the"),
				strings2.SingleCaseWord("good-doers"), // If Space wins, "good-doers" is one part. It is classified as SingleCaseWord if all lower chars?
				// "good-doers" -> IsAllLower? '-' is not letter.
				// loop: 'g' Lower, 'o' Lower... '-' !Letter (IsLower check passes as true if logic matches previous discussion?).
				// ClassifyPart:
				// if !unicode.IsLower(r) && unicode.IsLetter(r) { isAllLower = false }
				// '-' is not letter. So isAllLower remains true.
				// So it should be SingleCaseWord.
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
		{
			name:  "XMLReader SmartAcronyms",
			input: "XMLReader",
			opts:  []strings2.ParserOption{strings2.WithSmartAcronyms(true)},
			expected: []strings2.Word{
				strings2.AcronymWord("XML"),
				strings2.FirstUpperCaseWord("Reader"),
			},
		},
		{
			name:  "XMLReader No SmartAcronyms",
			input: "XMLReader",
			opts:  []strings2.ParserOption{strings2.WithSmartAcronyms(false)},
			expected: []strings2.Word{
				strings2.UpperCaseWord("XML"),
				strings2.FirstUpperCaseWord("Reader"),
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
