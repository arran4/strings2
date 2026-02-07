package strings2

import (
	"testing"
)

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
			got := upperCaseFirstLower(tt.input)
			if got != tt.expected {
				t.Errorf("upperCaseFirstLower(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUpperCaseFirstLower_Allocations(t *testing.T) {
	// Tests that no allocation occurs if the string is already correct
	input := "Test"
	if testing.AllocsPerRun(10, func() {
		upperCaseFirstLower(input)
	}) > 0 {
		t.Errorf("upperCaseFirstLower(%q) allocated memory when no change was needed", input)
	}

	// Test that allocation occurs when change IS needed
	input2 := "test"
	if testing.AllocsPerRun(10, func() {
		upperCaseFirstLower(input2)
	}) == 0 {
		t.Errorf("upperCaseFirstLower(%q) did not allocate memory when change was needed", input2)
	}
}
