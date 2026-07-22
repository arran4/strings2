package mappers_test

import (
	"strings"
	"testing"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func TestMapLowercase(t *testing.T) {
	mapper := mappers.MapLowercase("to", "at")
	input := []strings2.Word{
		strings2.SingleCaseWord("Went"),
		strings2.SingleCaseWord("TO"),
		strings2.SingleCaseWord("work"),
		strings2.SingleCaseWord("At"),
	}
	output := mapper(input)
	if len(output) != 4 {
		t.Fatalf("expected 4 words, got %d", len(output))
	}

	if output[1].String() != "to" {
		t.Errorf("expected to, got %s", output[1].String())
	}
	if output[3].String() != "at" {
		t.Errorf("expected at, got %s", output[3].String())
	}
	if _, ok := output[3].(strings2.SingleCaseWord); !ok {
		t.Errorf("expected SingleCaseWord for output[3], got %T", output[3])
	}
}

func TestMapLowercaseFromReader(t *testing.T) {
	reader := strings.NewReader("to\nat\n")
	mapper, err := mappers.MapLowercaseFromReader(reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := []strings2.Word{
		strings2.SingleCaseWord("Went"),
		strings2.SingleCaseWord("TO"),
		strings2.SingleCaseWord("work"),
		strings2.SingleCaseWord("At"),
	}
	output := mapper(input)

	if output[1].String() != "to" {
		t.Errorf("expected to, got %s", output[1].String())
	}
	if output[3].String() != "at" {
		t.Errorf("expected at, got %s", output[3].String())
	}
}
