package strings2

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		{
			name: "Basic",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			expected: "helloWorld",
		},
		{
			name: "With FirstUpperCaseWord",
			words: []Word{
				SingleCaseWord("hello"),
				FirstUpperCaseWord("world"),
			},
			expected: "helloWorld",
		},
		{
			name: "Single Word",
			words: []Word{
				SingleCaseWord("Hello"),
			},
			expected: "hello",
		},
		{
			name:     "Empty",
			words:    []Word{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCamelCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		{
			name: "Basic",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			expected: "HelloWorld",
		},
		{
			name: "With FirstUpperCaseWord",
			words: []Word{
				SingleCaseWord("hello"),
				FirstUpperCaseWord("world"),
			},
			expected: "HelloWorld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPascalCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		{
			name: "Basic",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			expected: "hello-world",
		},
		{
			name: "Screaming Kebab",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			options:  []Option{OptionCaseMode(CMScreaming)},
			expected: "HELLO-WORLD",
		},
		{
			name: "Double Delimiter",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			options:  []Option{OptionUpperIndicator("-")},
			expected: "hello--world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToKebabCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		{
			name: "Basic",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			expected: "hello_world",
		},
		{
			name: "Screaming Snake",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			options:  []Option{OptionCaseMode(CMScreaming)},
			expected: "HELLO_WORLD",
		},
		{
			name: "Using OptionFirstUpper",
			words: []Word{
				SingleCaseWord("hello"),
				SingleCaseWord("world"),
			},
			options:  []Option{OptionFirstUpper()},
			expected: "Hello_world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnakeCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToFormattedCase_CaseModes(t *testing.T) {
	words := []Word{
		SingleCaseWord("hello"),
		SingleCaseWord("world"),
	}

	tests := []struct {
		name     string
		caseMode CaseMode
		expected string
	}{
		{
			name:     "Verbatim",
			caseMode: CMVerbatim,
			expected: "hello-world", // Default delimiter is -
		},
		{
			name:     "Screaming",
			caseMode: CMScreaming,
			expected: "HELLO-WORLD",
		},
		{
			name:     "Whispering",
			caseMode: CMWhispering,
			expected: "hello-world",
		},
		{
			name:     "AllTitle",
			caseMode: CMAllTitle,
			expected: "Hello-World",
		},
		{
			name:     "FirstTitle",
			caseMode: CMFirstTitle,
			expected: "Hello-world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToFormattedCase(words, OptionCaseMode(tt.caseMode))
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestMixCaseSupport(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		{
			name: "MixCase Kebab",
			words: []Word{
				ExactCaseWord("camelCase"),
				SingleCaseWord("test"),
			},
			options:  []Option{OptionMixCaseSupport(), OptionDelimiter("-")},
			expected: "camel-Case-test",
		},
		{
			name: "FirstUpperCase MixCase",
			words: []Word{
				ExactCaseWord("camelCase"),
			},
			options:  []Option{OptionMixCaseSupport(), OptionDelimiter("_"), OptionFirstUpper()},
			expected: "Camel_Case",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToFormattedCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestUpperCaseFirst_Correctness(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "Test"},
		{"äpfel", "Äpfel"},
		{"ßeta", "ßeta"},
		{"", ""},
		{"a", "A"},
		{"A", "A"},
		{"1test", "1test"},
		{"!test", "!test"},
		{"Test", "Test"},
		{"Öpfel", "Öpfel"},
		{"\xff", "\xff"}, // Invalid UTF-8
		{"\xe2\x82\x28", "\xe2\x82\x28"}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		got := UpperCaseFirst(tt.input)
		if got != tt.expected {
			t.Errorf("UpperCaseFirst(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
