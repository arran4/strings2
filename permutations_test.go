package strings2_test

import (
	"testing"
	"github.com/arran4/strings2"
)

func TestToTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic", "the lord of the rings", "The Lord of the Rings"},
		{"screaming snake", "A_NEW_HOPE", "A New Hope"},
		{"camel case", "camelCaseInput", "Camel Case Input"},
		{"mixed kebab", "mixed-UP-Kebab", "Mixed UP Kebab"},
		{"prepositions", "this is a test of the smart title case", "This Is a Test of the Smart Title Case"},
		{"first word prep", "of mice and men", "Of Mice and Men"},
		{"last word prep", "who are you looking at", "Who Are You Looking At"},
		{"acronym request", "parse HTTP request", "Parse HTTP Request"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := strings2.ToTitle(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
