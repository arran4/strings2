package strings2_test

import (
	"strings"
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
			formatDelim:      "_", // SnakeCase formatter uses "_" implicitly, but we might check if we can pass it explicitly or if defaults matter.
            // ToSnakeCase default delim is "_".
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
				restored = strings2.ToFormattedCase(words, strings2.OptionDelimiter(tc.formatDelim))
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
