package strings2

import (
	"reflect"
	"strings"
	"testing"
)

// TestOptionMetadata verifies that options correctly modify the internal caseConfig metadata.
// This is "meta data testing" as it verifies the configuration state rather than just the output.
func TestOptionMetadata(t *testing.T) {
	tests := []struct {
		name     string
		options  []Option
		expected caseConfig
	}{
		{
			name:    "Default Config",
			options: []Option{},
			expected: caseConfig{
				delimiter: "-", // Default is "-"
			},
		},
		{
			name: "Set Delimiter",
			options: []Option{
				OptionDelimiter("_"),
			},
			expected: caseConfig{
				delimiter: "_",
			},
		},
		{
			name: "Case Mode Screaming",
			options: []Option{
				OptionCaseMode(CMScreaming),
			},
			expected: caseConfig{
				caseMode:  CMScreaming,
				delimiter: "-",
			},
		},
		{
			name: "First Upper",
			options: []Option{
				OptionFirstUpper(),
			},
			expected: caseConfig{
				firstUpper: true,
				delimiter:  "-",
			},
		},
		{
			name: "First Lower",
			options: []Option{
				OptionFirstLower(),
			},
			expected: caseConfig{
				firstLower: true,
				delimiter:  "-",
			},
		},
		{
			name: "Mix Case Support",
			options: []Option{
				OptionMixCaseSupport(),
			},
			expected: caseConfig{
				mixCaseSupport: true,
				delimiter:      "-",
			},
		},
		{
			name: "Upper Indicator",
			options: []Option{
				OptionUpperIndicator("="),
			},
			expected: caseConfig{
				upperIndicator: "=",
				delimiter:      "-",
			},
		},
		{
			name: "Combination",
			options: []Option{
				OptionDelimiter("."),
				OptionCaseMode(CMWhispering),
				OptionFirstUpper(),
			},
			expected: caseConfig{
				delimiter:  ".",
				caseMode:   CMWhispering,
				firstUpper: true,
			},
		},
		{
			name: "Last Write Wins (Delimiter)",
			options: []Option{
				OptionDelimiter("-"),
				OptionDelimiter("_"),
			},
			expected: caseConfig{
				delimiter: "_",
			},
		},
		{
			name: "Last Write Wins (CaseMode)",
			options: []Option{
				OptionCaseMode(CMFirstTitle),
				OptionCaseMode(CMScreaming),
			},
			expected: caseConfig{
				caseMode:  CMScreaming,
				delimiter: "-",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize with defaults as ToFormattedCase does
			cfg := &caseConfig{delimiter: "-"}
			for _, opt := range tt.options {
				opt(cfg)
			}

			if !reflect.DeepEqual(*cfg, tt.expected) {
				t.Errorf("metadata mismatch:\ngot  %+v\nwant %+v", *cfg, tt.expected)
			}
		})
	}
}

// TestCircularConsistency verifies that splitting a formatted string and re-formatting it
// with the same options results in the same string. This ensures the operations are circular
// (idempotent/stable) for ExactCaseWord.
func TestCircularConsistency(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter string
		options   []Option
	}{
		{
			name:      "Simple Kebab",
			input:     "hello-world",
			delimiter: "-",
			options:   []Option{OptionDelimiter("-")},
		},
		{
			name:      "Snake Case",
			input:     "hello_world",
			delimiter: "_",
			options:   []Option{OptionDelimiter("_")},
		},
		{
			name:      "Screaming Snake",
			input:     "HELLO_WORLD",
			delimiter: "_",
			options:   []Option{OptionDelimiter("_"), OptionCaseMode(CMScreaming)},
		},
		{
			name:      "Pascal Case (treated as exact words joined by empty)",
			input:     "HelloWorld",
			delimiter: "",
			options:   []Option{OptionDelimiter(""), OptionCaseMode(CMAllTitle)},
			// Note: Splitting "HelloWorld" by "" gives individual characters, which isn't what we want usually.
			// However, for this test, if we split by expected delimiter...
			// If delimiter is empty string, strings.Split("abc", "") -> ["", "a", "b", "c", ""].
			// This test logic needs to handle empty delimiter carefully or skip it.
			// Let's stick to non-empty delimiters for simple circular tests or handle it specifically.
		},
		{
			name:      "Double Delimiter",
			input:     "hello--world",
			delimiter: "--", // Effectively splitting by the double delimiter
			options:   []Option{OptionDelimiter("-"), OptionUpperIndicator("-")},
			// Wait, OptionUpperIndicator("-") with Delimiter("-") makes delimiter "--".
		},
		{
			name:      "Three Words",
			input:     "one-two-three",
			delimiter: "-",
			options:   []Option{OptionDelimiter("-")},
		},
		{
			name:      "Already Mixed Case Preserved",
			input:     "camel-Case",
			delimiter: "-",
			options:   []Option{OptionDelimiter("-")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.delimiter == "" {
				t.Skip("Skipping empty delimiter circular test as split behavior is char-based")
			}

			// 1. Split the string
			parts := strings.Split(tt.input, tt.delimiter)

			// 2. Wrap in ExactCaseWord
			words := make([]Word, len(parts))
			for i, p := range parts {
				words[i] = ExactCaseWord(p)
			}

			// 3. Re-format
			result := ToFormattedCase(words, tt.options...)

			// 4. Verify
			if result != tt.input {
				t.Errorf("Circular test failed: split %q -> reformatted %q", tt.input, result)
			}
		})
	}
}

// TestInternalFlagsConsistency ensures that options set flags that are internally consistent
// or that conflicting flags are handled deterministically (Meta Data logic).
func TestInternalFlagsConsistency(t *testing.T) {
    // This tests the logic inside ToFormattedCase that derives internal flags from CaseMode
	// We can't access local variables inside ToFormattedCase, but we can verify the outcome
	// matches the expected behavior of those flags.

    // Case 1: CMScreaming sets cfg.screaming = true.
    // SingleCaseWord should be Uppercased.
    t.Run("CMScreaming implies screaming", func(t *testing.T) {
        res := ToFormattedCase([]Word{SingleCaseWord("hello")}, OptionCaseMode(CMScreaming))
        if res != "HELLO" {
             t.Errorf("CMScreaming did not result in screaming output: %q", res)
        }
    })

    // Case 2: CMWhispering implies whispering.
    // SingleCaseWord should be Lowercased (even if input is upper).
    t.Run("CMWhispering implies whispering", func(t *testing.T) {
        res := ToFormattedCase([]Word{SingleCaseWord("HELLO")}, OptionCaseMode(CMWhispering))
        if res != "hello" {
             t.Errorf("CMWhispering did not result in whispering output: %q", res)
        }
    })
}
