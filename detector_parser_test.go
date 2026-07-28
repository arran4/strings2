package strings2

import (
	"testing"
)

func TestParseDelimiterDetector(t *testing.T) {
	expr := "any(numeric, whitespace, nonalphanumeric)"
	det, err := ParseDelimiterDetector(expr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	subs, _ := StringToSubParts("A1 -")

	if l := det(subs, 0); l != 0 {
		t.Errorf("expected 0 for A, got %d", l)
	}
	if l := det(subs, 1); l != 1 {
		t.Errorf("expected 1 for 1, got %d", l)
	}
	if l := det(subs, 2); l != 1 {
		t.Errorf("expected 1 for space, got %d", l)
	}
	if l := det(subs, 3); l != 1 {
		t.Errorf("expected 1 for -, got %d", l)
	}
}
