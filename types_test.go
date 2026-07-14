package strings2

import (
	"errors"
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
			result, _ := ToCamelCase(tt.words, tt.options...)
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
			result, _ := ToPascalCase(tt.words, tt.options...)
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
			options:  []Option{OptionScreaming()},
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
			result, _ := ToKebabCase(tt.words, tt.options...)
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
			options:  []Option{OptionScreaming()},
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
			result, _ := ToSnakeCase(tt.words, tt.options...)
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
		{"\xff", "\xff"},                 // Invalid UTF-8
		{"\xe2\x82\x28", "\xe2\x82\x28"}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		got := UpperCaseFirst(tt.input)
		if got != tt.expected {
			t.Errorf("UpperCaseFirst(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestUpperCaseFirstWithErr_Correctness(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		{"test", "Test", false},
		{"äpfel", "Äpfel", false},
		{"", "", false},
		{"\xff", "", true},         // Invalid UTF-8
		{"\xe2\x82\x28", "", true}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		got, err := UpperCaseFirstWithErr(tt.input)
		if tt.expectError {
			if err == nil {
				t.Errorf("UpperCaseFirstWithErr(%q) expected error, got nil", tt.input)
			}
			if !errors.Is(err, ErrRune) {
				t.Errorf("UpperCaseFirstWithErr(%q) expected ErrRune, got %v", tt.input, err)
			}
		} else {
			if err != nil {
				t.Errorf("UpperCaseFirstWithErr(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("UpperCaseFirstWithErr(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		}
	}
}

func TestMustUpperCaseFirst_Correctness(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectPanic bool
	}{
		{"test", "Test", false},
		{"\xff", "", true},         // Invalid UTF-8
		{"\xe2\x82\x28", "", true}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("MustUpperCaseFirst(%q) panicked unexpectedly: %v", tt.input, r)
					}
				} else {
					if tt.expectPanic {
						t.Errorf("MustUpperCaseFirst(%q) expected panic, but did not panic", tt.input)
					}
				}
			}()
			got := MustUpperCaseFirst(tt.input)
			if !tt.expectPanic && got != tt.expected {
				t.Errorf("MustUpperCaseFirst(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		}()
	}
}

func TestLowerCaseFirst_Correctness(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Test", "test"},
		{"Äpfel", "äpfel"},
		{"ßeta", "ßeta"}, // ß doesn't change
		{"", ""},
		{"A", "a"},
		{"a", "a"},
		{"1test", "1test"},
		{"!test", "!test"},
		{"test", "test"},
		{"Öpfel", "öpfel"},
		{"\xff", "\xff"},                 // Invalid UTF-8
		{"\xe2\x82\x28", "\xe2\x82\x28"}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		got := LowerCaseFirst(tt.input)
		if got != tt.expected {
			t.Errorf("LowerCaseFirst(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestLowerCaseFirstWithErr_Correctness(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		{"Test", "test", false},
		{"Äpfel", "äpfel", false},
		{"", "", false},
		{"\xff", "", true},         // Invalid UTF-8
		{"\xe2\x82\x28", "", true}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		got, err := LowerCaseFirstWithErr(tt.input)
		if tt.expectError {
			if err == nil {
				t.Errorf("LowerCaseFirstWithErr(%q) expected error, got nil", tt.input)
			}
			if !errors.Is(err, ErrRune) {
				t.Errorf("LowerCaseFirstWithErr(%q) expected ErrRune, got %v", tt.input, err)
			}
		} else {
			if err != nil {
				t.Errorf("LowerCaseFirstWithErr(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("LowerCaseFirstWithErr(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		}
	}
}

func TestMustLowerCaseFirst_Correctness(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectPanic bool
	}{
		{"Test", "test", false},
		{"\xff", "", true},         // Invalid UTF-8
		{"\xe2\x82\x28", "", true}, // Invalid UTF-8 sequence
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("MustLowerCaseFirst(%q) panicked unexpectedly: %v", tt.input, r)
					}
				} else {
					if tt.expectPanic {
						t.Errorf("MustLowerCaseFirst(%q) expected panic, but did not panic", tt.input)
					}
				}
			}()
			got := MustLowerCaseFirst(tt.input)
			if !tt.expectPanic && got != tt.expected {
				t.Errorf("MustLowerCaseFirst(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		}()
	}
}

func TestToFormattedCase_MultibyteFirstLower(t *testing.T) {
	// Tests that OptionFirstLower correctly handles multibyte characters.
	words := []Word{ExactCaseWord("Äpfel")}
	got := ToFormattedCase(words, OptionFirstLower())
	want := "äpfel"
	if got != want {
		t.Errorf("ToFormattedCase with OptionFirstLower for %q = %q, want %q", "Äpfel", got, want)
	}
}

func TestOptionUTF8Modes(t *testing.T) {
	tests := []struct {
		name      string
		words     []Word
		options   []Option
		expectErr bool
		expected  string
	}{
		{
			name: "Strict Mode Error",
			words: []Word{
				FirstUpperCaseWord("\xfftest"),
			},
			options:   []Option{OptionStrict()},
			expectErr: true,
		},
		{
			name: "Loose Mode Preserves Invalid",
			words: []Word{
				FirstUpperCaseWord("\xfftest"),
			},
			options:   []Option{OptionLoose()},
			expectErr: false,
			expected:  "\xfftest",
		},
		{
			name: "Default Mode Replaces Invalid",
			words: []Word{
				FirstUpperCaseWord("\xfftest"),
			},
			options:   []Option{}, // Default is UTF8Replace
			expectErr: false,
			expected:  "\uFFFDtest",
		},
		{
			name: "SingleCaseWord CMAllTitle Strict",
			words: []Word{
				SingleCaseWord("\xfftest"),
			},
			options:   []Option{OptionCaseMode(CMAllTitle), OptionStrict()},
			expectErr: true,
			expected:  "",
		},
		{
			name: "SingleCaseWord CMAllTitle Loose",
			words: []Word{
				SingleCaseWord("\xfftest"),
			},
			options:   []Option{OptionCaseMode(CMAllTitle), OptionLoose()},
			expectErr: false,
			expected:  "\xfftest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WordsToFormattedCase(tt.words, convertOptions(tt.options)...)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if !errors.Is(err, ErrRune) {
					t.Errorf("expected ErrRune, got %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if got != tt.expected {
					t.Errorf("got %q (bytes: %x), want %q (bytes: %x)", got, []byte(got), tt.expected, []byte(tt.expected))
				}
			}
		})
	}
}

func TestUpperCaseWord_Verbatim_Bug(t *testing.T) {
	// "HELLO" is parsed as AcronymWord by default (SmartAcronyms=true).
	// But if SmartAcronyms=false, it becomes UpperCaseWord.

	input := "HELLO"

	// Case 1: Default (SmartAcronyms=true)
	words1, _ := Parse(input)      // [AcronymWord("HELLO")]
	res1, _ := ToSnakeCase(words1) // ToSnakeCase defaults to Verbatim (but with delimiter "_")
	// AcronymWord preserves case by default.
	if res1 != "HELLO" {
		t.Errorf("Default behavior changed? Got %q, want %q", res1, "HELLO")
	}

	// Case 2: SmartAcronyms=false
	words2, _ := Parse(input, WithSmartAcronyms(false)) // [UpperCaseWord("HELLO")]
	// Expectation: Verbatim mode should preserve case -> "HELLO"
	res2, _ := ToSnakeCase(words2)

	expected := "HELLO"
	if res2 != expected {
		t.Errorf("UpperCaseWord (SmartAcronyms=false) did not preserve case. Got %q, want %q", res2, expected)
	}
}

func TestUpperCaseFirstLower_Correctness(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: "",
		},
		{
			name:     "ASCII Lower",
			input:    "test",
			expected: "Test",
		},
		{
			name:     "ASCII Mixed",
			input:    "tEsT",
			expected: "Test",
		},
		{
			name:     "ASCII Upper",
			input:    "TEST",
			expected: "Test",
		},
		{
			name:     "Already Correct",
			input:    "Test",
			expected: "Test",
		},
		{
			name:     "Unicode Lower",
			input:    "äpfel",
			expected: "Äpfel",
		},
		{
			name:     "Unicode Upper",
			input:    "ÄPFEL",
			expected: "Äpfel",
		},
		{
			name:     "Unicode Mixed",
			input:    "äPfEl",
			expected: "Äpfel",
		},
		{
			name:     "Special Char Start",
			input:    "!test",
			expected: "!test",
		},
		{
			name:     "Number Start",
			input:    "1test",
			expected: "1test",
		},
		{
			name:     "Invalid UTF-8",
			input:    "\xff\xfe\xfd",
			expected: "\uFFFD\uFFFD\uFFFD",
		},
		{
			name:     "Partial Invalid UTF-8",
			input:    "test\xff",
			expected: "Test\uFFFD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := upperCaseFirstLower(tt.input, UTF8Replace)
			if err != nil {
				t.Errorf("upperCaseFirstLower(%q, UTF8Replace) returned unexpected error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("upperCaseFirstLower(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUpperCaseFirstLower_Strict(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "Valid ASCII",
			input:     "test",
			expectErr: false,
		},
		{
			name:      "Valid Unicode",
			input:     "äpfel",
			expectErr: false,
		},
		{
			name:      "Invalid UTF-8 Start",
			input:     "\xfftest",
			expectErr: true,
		},
		{
			name:      "Invalid UTF-8 Middle",
			input:     "te\xffst",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := upperCaseFirstLower(tt.input, UTF8Strict)
			if tt.expectErr {
				if err == nil {
					t.Errorf("upperCaseFirstLower(%q, UTF8Strict) expected error, got nil", tt.input)
				}
				if !errors.Is(err, ErrRune) {
					t.Errorf("upperCaseFirstLower(%q, UTF8Strict) expected ErrRune, got %v", tt.input, err)
				}
			} else {
				if err != nil {
					t.Errorf("upperCaseFirstLower(%q, UTF8Strict) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

func TestUpperCaseFirstLower_Loose(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Invalid UTF-8 Start",
			input:    "\xfftest",
			expected: "\xfftest", // Preserves invalid byte
		},
		{
			name:     "Invalid UTF-8 Middle",
			input:    "te\xffst",
			expected: "Te\xffst", // Preserves invalid byte, title cases valid parts
		},
		{
			name:     "Mixed Invalid",
			input:    "\xffT\xff",
			expected: "\xfft\xff", // Start invalid kept, 'T' -> 't', 't' lowercased? No wait.
			// upperCaseFirstLower Logic:
			// 1. Decode first rune. If invalid: write byte.
			// 2. Loop rest. If invalid: write byte. Else toLower.
			// Input: \xff T \xff
			// 1. First: \xff. Invalid. Write \xff.
			// 2. Rest: "T\xff".
			//    - 'T': ToLower -> 't'.
			//    - \xff: Invalid. Write \xff.
			// Result: "\xfft\xff".
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := upperCaseFirstLower(tt.input, UTF8Ignore)
			if err != nil {
				t.Errorf("upperCaseFirstLower(%q, UTF8Ignore) returned unexpected error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("upperCaseFirstLower(%q, UTF8Ignore) = %q (bytes: %x), want %q (bytes: %x)", tt.input, got, []byte(got), tt.expected, []byte(tt.expected))
			}
		})
	}
}

func TestUpperCaseFirstLower_Allocations(t *testing.T) {
	// Tests that no allocation occurs if the string is already correct
	input := "Test"
	if testing.AllocsPerRun(10, func() {
		_, _ = upperCaseFirstLower(input, UTF8Replace)
	}) > 0 {
		t.Errorf("upperCaseFirstLower(%q) allocated memory when no change was needed", input)
	}

	// Test that allocation occurs when change IS needed
	input2 := "test"
	if testing.AllocsPerRun(10, func() {
		_, _ = upperCaseFirstLower(input2, UTF8Replace)
	}) == 0 {
		t.Errorf("upperCaseFirstLower(%q) did not allocate memory when change was needed", input2)
	}
}
