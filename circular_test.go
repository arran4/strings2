package strings2_test

import (
	"testing"

	"github.com/arran4/strings2"
)

func TestCircularRestoration(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedOverride *string // If set, expects this output instead of input (lossy)
		partitionerCfg   strings2.PartitionerConfig
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
			name:  "Lossy conversion (eats separators)",
			input: "Hello-world",
			// No PreserveSep, so '-' is lost
			partitionerCfg: strings2.PartitionerConfig{
				Delimiters: map[rune]bool{'-': true},
			},
			formatDelim: "",
			expectedOverride: func() *string {
				s := "Helloworld"
				return &s
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			words, err := strings2.Parse(tc.input, tc.partitionerCfg)
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}

			restored := strings2.ToFormattedCase(words, strings2.OptionDelimiter(tc.formatDelim))

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
