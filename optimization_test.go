package strings2

import (
	"testing"
)

func TestOptimizationCorrectness(t *testing.T) {
	// This test specifically targets the SingleCaseWord optimization path
	// to ensure that mixed-case inputs are handled identically to the original
	// implementation (which normalized to lowercase via .String() first).

	tests := []struct {
		name     string
		input    []Word
		options  []Option
		expected string
	}{
		{
			name:     "MixedCase Default (Implied Lower)",
			input:    []Word{SingleCaseWord("HeLLo"), SingleCaseWord("WoRLd")},
			options:  []Option{OptionDelimiter("-")},
			expected: "hello-world",
		},
		{
			name:     "MixedCase Screaming",
			input:    []Word{SingleCaseWord("HeLLo"), SingleCaseWord("WoRLd")},
			options:  []Option{OptionDelimiter("-"), OptionCaseMode(CMScreaming)},
			expected: "HELLO-WORLD",
		},
		{
			name:     "MixedCase Whispering",
			input:    []Word{SingleCaseWord("HeLLo"), SingleCaseWord("WoRLd")},
			options:  []Option{OptionDelimiter("-"), OptionCaseMode(CMWhispering)},
			expected: "hello-world",
		},
		{
			name:     "MixedCase AllTitle",
			input:    []Word{SingleCaseWord("HeLLo"), SingleCaseWord("WoRLd")},
			options:  []Option{OptionDelimiter("-"), OptionCaseMode(CMAllTitle)},
			expected: "Hello-World",
		},
		{
			name:     "MixedCase FirstTitle",
			input:    []Word{SingleCaseWord("HeLLo"), SingleCaseWord("WoRLd")},
			options:  []Option{OptionDelimiter("-"), OptionCaseMode(CMFirstTitle)},
			expected: "Hello-world",
		},
		// Edge cases
		{
			name:     "Empty String",
			input:    []Word{SingleCaseWord("")},
			options:  []Option{OptionDelimiter("-")},
			expected: "",
		},
		{
			name:     "All Caps Input Default",
			input:    []Word{SingleCaseWord("HELLO")},
			options:  []Option{OptionDelimiter("-")},
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToFormattedCase(tt.input, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
