package strings2

import (
	"testing"
)

func TestParseDelimiterDetector(t *testing.T) {
	tests := []struct {
		expr string
		subs string
		expected []int
	}{
		{
			expr: "any(numeric, whitespace, nonalphanumeric)",
			subs: "A1 -",
			expected: []int{0, 1, 1, 1},
		},
		{
			expr: "s(\"_*-\")",
			subs: "A_*- ",
			expected: []int{0, 1, 1, 1, 0},
		},
		{
			expr: "union(s(\"a\"), s(\"b\"))",
			subs: "abc",
			expected: []int{1, 1, 0},
		},
		{
			expr: "not(s(\" \"))",
			subs: "a b",
			expected: []int{1, 0, 1},
		},
	}

	for _, tt := range tests {
		det, err := ParseDelimiterDetector(tt.expr)
		if err != nil {
			t.Fatalf("unexpected error parsing %q: %v", tt.expr, err)
		}

		subs, _ := StringToSubParts(tt.subs)
		if len(subs) != len(tt.expected) {
			t.Fatalf("test case mismatch for %q: subs len %d != expected len %d", tt.expr, len(subs), len(tt.expected))
		}

		for i, expected := range tt.expected {
			if l := det(subs, i); l != expected {
				t.Errorf("expr %q, char %q at index %d: expected %d, got %d", tt.expr, subs[i].Rune(), i, expected, l)
			}
		}
	}
}
