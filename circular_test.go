package strings2_test

import (
	"testing"

	"github.com/arran4/strings2"
)

func TestCircularRestoration(t *testing.T) {
	s := "Helloworld"
	tests := []struct {
		name             string
		input            string
		expectedOverride *string // If set, expects this output instead of input (lossy)
		partitionerCfg   strings2.PartitionerConfig
		formatFunc       func([]strings2.Word, ...strings2.Option) string // Allow custom formatter
		formatDelim      string
	}{
		{
			name:  "Sentence with spaces",
			input: "Hello to all the good-doers out there",
			partitionerCfg: strings2.PartitionerConfig{
				Delimiters:  map[rune]bool{' ': true, '-': true},
				PreserveSep: true,
			},
			formatDelim: "",
		},
		{
			name:  "Complex delimiters",
			input: "user_name.with-mixed@delimiters",
			partitionerCfg: strings2.PartitionerConfig{
				Delimiters:  map[rune]bool{'_': true, '.': true, '-': true, '@': true},
				PreserveSep: true,
			},
			formatDelim: "",
		},
		{
			name:  "Consecutive delimiters",
			input: "item1,,item2--item3",
			partitionerCfg: strings2.PartitionerConfig{
				Delimiters:  map[rune]bool{',': true, '-': true},
				PreserveSep: true,
			},
			formatDelim: "",
		},
		{
			name:  "Lossy conversion (eats separators)",
			input: "Hello-world",
			// No PreserveSep, so '-' is lost
			partitionerCfg: strings2.PartitionerConfig{
				Delimiters: map[rune]bool{'-': true},
			},
			formatDelim:      "",
			expectedOverride: &s,
		},
		{
			name:  "Snake to Snake Roundtrip",
			input: "hello_world_test",
			partitionerCfg: strings2.PartitionerConfig{
				Delimiters: map[rune]bool{'_': true},
			},
			formatFunc: func(words []strings2.Word, opts ...strings2.Option) string {
				return strings2.ToSnakeCase(words, opts...)
			},
			formatDelim:      "_",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			words, err := strings2.Parse(tc.input, tc.partitionerCfg)
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}

			var restored string
			if tc.formatFunc != nil {
				restored = tc.formatFunc(words)
			} else {
				// Use the new generic WordsToFormattedCase
				// Assuming WordsToFormattedCase is what ToFormattedCase was doing but error aware
				// Since ToFormattedCase is deprecated, we should use WordsToFormattedCase if we can,
				// or ToFormattedCase for backward compat in this test context?
				// The test logic here uses ToFormattedCase wrapper from earlier.
				// Let's use WordsToFormattedCase and ignore error for this test as we don't expect formatting errors here.

				// Need to convert OptionDelimiter to ...any or ...Option
				// WordsToFormattedCase accepts ...any

				res, _ := strings2.WordsToFormattedCase(words, strings2.OptionDelimiter(tc.formatDelim))
				restored = res
			}

			expected := tc.input
			if tc.expectedOverride != nil {
				expected = *tc.expectedOverride
			}

			if restored != expected {
				t.Errorf("Circular restoration failed.\nInput:    %q\nExpected: %q\nGot:      %q", tc.input, expected, restored)
			}
		})
	}
}
