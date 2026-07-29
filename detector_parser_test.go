package strings2

import (
	"testing"
)

func TestParseDelimiterDetector(t *testing.T) {
	tests := []struct {
		expr     string
		subs     string
		expected []int
		wantErr  bool
	}{
		{
			expr:     "any(numeric, whitespace, nonalphanumeric)",
			subs:     "A1 -",
			expected: []int{0, 1, 1, 1},
		},
		{
			expr:     "s(\"_*-\")",
			subs:     "A_*- ",
			expected: []int{0, 1, 1, 1, 0},
		},
		{
			expr:     "union(s(\"a\"), s(\"b\"))",
			subs:     "abc",
			expected: []int{1, 1, 0},
		},
		{
			expr:     "not(s(\" \"))",
			subs:     "a b",
			expected: []int{1, 0, 1},
		},
		{
			expr:     "delimiters(\"xy\")",
			subs:     "axyz",
			expected: []int{0, 1, 1, 0},
		},
		{
			expr:     "tab",
			subs:     "a\tb",
			expected: []int{0, 1, 0},
		},
		{
			expr:     "whitespace",
			subs:     "a \n\t\rb",
			expected: []int{0, 1, 1, 1, 1, 0},
		},
		{
			expr:     "numeric",
			subs:     "a1b",
			expected: []int{0, 1, 0},
		},
		{
			expr:     "nonalphanumeric",
			subs:     "a1!@",
			expected: []int{0, 0, 1, 1},
		},
		{
			expr:     "any(numeric)",
			subs:     "a1",
			expected: []int{0, 1},
		},
		{
			expr:     "s(\"a\", \"b\")", // Too many args
			wantErr:  true,
		},
		{
			expr:     "not(numeric, whitespace)", // Too many args
			wantErr:  true,
		},
		{
			expr:     "unknown(numeric)", // Unknown func
			wantErr:  true,
		},
		{
			expr:     "unknown_ident", // Unknown ident
			wantErr:  true,
		},
		{
			expr:     "s(\"\\\"\")", // Escaped quote
			subs:     "a\"b",
			expected: []int{0, 1, 0},
		},
		{
			expr:     "s(\"\\\\\")", // Escaped backslash
			subs:     "a\\b",
			expected: []int{0, 1, 0},
		},
		{
			expr:     "s(\"unterminated",
			wantErr:  true,
		},
		{
			expr:     "not(",
			wantErr:  true,
		},
		{
			expr:     "any(numeric))",
			wantErr:  true,
		},
		{
			expr:     "  s  (  \" a \"  )  ",
			subs:     "b a ",
			expected: []int{0, 1, 1, 1},
		},
		{
			expr:     "not(any(numeric, whitespace))",
			subs:     "a1 ",
			expected: []int{1, 0, 0},
		},
	}

	for _, tt := range tests {
		det, err := ParseDelimiterDetector(tt.expr)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseDelimiterDetector(%q) error = %v, wantErr %v", tt.expr, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			continue
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
