package strings2_test

import (
	"testing"

	"github.com/arran4/strings2"
)

func TestSubPartsToParts(t *testing.T) {
	input := "hello_world"
	subs, _ := strings2.StringToSubParts(input)

	t.Run("SnakeCasePartitioner", func(t *testing.T) {
		parts := strings2.SubPartsToParts(subs, strings2.SnakeCasePartitioner)
		if len(parts) != 2 {
			t.Errorf("Expected 2 parts, got %d", len(parts))
		}
		if parts[0].String() != "hello" {
			t.Errorf("Part 0 mismatch: %s", parts[0].String())
		}
		if parts[1].String() != "world" {
			t.Errorf("Part 1 mismatch: %s", parts[1].String())
		}
	})
}

func TestCamelCasePartitioner(t *testing.T) {
	input := "PDFLoader"
	subs, _ := strings2.StringToSubParts(input)

	parts := strings2.SubPartsToParts(subs, strings2.CamelCasePartitioner)
	if len(parts) != 2 {
		t.Fatalf("Expected 2 parts, got %d", len(parts))
	}
	if parts[0].String() != "PDF" {
		t.Errorf("Part 0 mismatch: %s", parts[0].String())
	}
	if parts[1].String() != "Loader" {
		t.Errorf("Part 1 mismatch: %s", parts[1].String())
	}
}
