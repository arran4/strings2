package strings2_test

import (
	"testing"

	"github.com/arran4/strings2"
)

func TestStringToSubParts(t *testing.T) {
	input := "Hi_123"
	subs, stats := strings2.StringToSubParts(input)

	if len(subs) != 6 {
		t.Errorf("Expected 6 subparts, got %d", len(subs))
	}

	if subs[0].Rune() != 'H' || !subs[0].IsUpper() {
		t.Errorf("First subpart mismatch")
	}
	if subs[2].Rune() != '_' || !subs[2].IsSymbol() {
		t.Errorf("Third subpart mismatch (symbol)")
	}
	if subs[3].Rune() != '1' || !subs[3].IsDigit() {
		t.Errorf("Fourth subpart mismatch (digit)")
	}

	if stats.Upper != 1 {
		t.Errorf("Stats Upper mismatch: %d", stats.Upper)
	}
	if stats.Digits != 3 {
		t.Errorf("Stats Digits mismatch: %d", stats.Digits)
	}
	if stats.SymbolCounts['_'] != 1 {
		t.Errorf("Stats SymbolCounts mismatch")
	}
}
