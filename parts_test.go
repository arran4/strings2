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

func TestDelimiterDetector(t *testing.T) {
	input := "foo::bar--baz"
	subs, _ := strings2.StringToSubParts(input)

	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		DelimiterDetector: func(subs []strings2.SubPart, index int) int {
			if index+1 < len(subs) && subs[index].Rune() == ':' && subs[index+1].Rune() == ':' {
				return 2
			}
			if index+1 < len(subs) && subs[index].Rune() == '-' && subs[index+1].Rune() == '-' {
				return 2
			}
			return 0
		},
		PreserveSep: true,
	})

	parts := strings2.SubPartsToParts(subs, partitioner)
	if len(parts) != 5 {
		t.Fatalf("Expected 5 parts, got %d", len(parts))
	}
	if parts[0].String() != "foo" {
		t.Errorf("Part 0 mismatch: %s", parts[0].String())
	}
	if parts[1].String() != "::" {
		t.Errorf("Part 1 mismatch: %s", parts[1].String())
	}
	if parts[2].String() != "bar" {
		t.Errorf("Part 2 mismatch: %s", parts[2].String())
	}
	if parts[3].String() != "--" {
		t.Errorf("Part 3 mismatch: %s", parts[3].String())
	}
	if parts[4].String() != "baz" {
		t.Errorf("Part 4 mismatch: %s", parts[4].String())
	}
}

func TestDelimiterDetectorOutOfBounds(t *testing.T) {
	input := "foo"
	subs, _ := strings2.StringToSubParts(input)

	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		DelimiterDetector: func(subs []strings2.SubPart, index int) int {
			// Malicious detector that returns a length larger than the slice
			return 100
		},
		PreserveSep: true,
	})

	// Should not panic, but return separator of constrained length
	parts := strings2.SubPartsToParts(subs, partitioner)
	if len(parts) != 1 {
		t.Fatalf("Expected 1 part, got %d", len(parts))
	}
	if parts[0].String() != "foo" {
		t.Errorf("Part 0 mismatch: %s", parts[0].String())
	}
	// "foo" is returned as a SeparatorPart because the detector claims it's a delimiter,
	// and the length is constrained to len(subs) - i = 3
	if _, ok := parts[0].(*strings2.SeparatorPart); !ok {
		t.Errorf("Expected SeparatorPart for malicious length, got %T", parts[0])
	}
}

func TestDelimiterDetectorNegativeBounds(t *testing.T) {
	input := "foo"
	subs, _ := strings2.StringToSubParts(input)

	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		DelimiterDetector: func(subs []strings2.SubPart, index int) int {
			return -100
		},
	})

	// Should not panic, should act as length 0
	parts := strings2.SubPartsToParts(subs, partitioner)
	if len(parts) != 1 {
		t.Fatalf("Expected 1 part, got %d", len(parts))
	}
	if parts[0].String() != "foo" {
		t.Errorf("Part 0 mismatch: %s", parts[0].String())
	}
}
