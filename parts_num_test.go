package strings2

import (
	"reflect"
	"testing"
)

func TestNumberMode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		mode     NumberMode
		expected []string
	}{
		// None
		{"None_User123ID", "User123ID", NumberModeNone, []string{"User123ID"}},
		{"None_UPPER123", "UPPER123", NumberModeNone, []string{"UPPER123"}},
		{"None_123test", "123test", NumberModeNone, []string{"123test"}},

		// SplitAlways
		{"SplitAlways_User123ID", "User123ID", NumberModeSplitAlways, []string{"User", "123", "ID"}},
		{"SplitAlways_UPPER123", "UPPER123", NumberModeSplitAlways, []string{"UPPER", "123"}},
		{"SplitAlways_123test", "123test", NumberModeSplitAlways, []string{"123", "test"}},

		// MergeWithWord
		{"MergeWithWord_User123ID", "User123ID", NumberModeMergeWithWord, []string{"User123", "ID"}},
		{"MergeWithWord_UPPER123", "UPPER123", NumberModeMergeWithWord, []string{"UPPER123"}},
		{"MergeWithWord_123test", "123test", NumberModeMergeWithWord, []string{"123test"}},

		// TreatAsLowercase
		{"TreatAsLowercase_User123ID", "User123ID", NumberModeTreatAsLowercase, []string{"User123", "ID"}},
		{"TreatAsLowercase_UPPER123", "UPPER123", NumberModeTreatAsLowercase, []string{"UPPE", "R123"}},
		{"TreatAsLowercase_123test", "123test", NumberModeTreatAsLowercase, []string{"123test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			words, err := Parse(tt.input, WithNumberMode(tt.mode))
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}
			var got []string
			for _, w := range words {
				got = append(got, w.String())
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse(%q) with mode %v = %v; want %v", tt.input, tt.mode, got, tt.expected)
			}
		})
	}
}
