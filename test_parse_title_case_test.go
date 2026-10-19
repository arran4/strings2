package strings2_test

import (
	"fmt"
	"testing"

	"github.com/arran4/strings2"
)

func TestParseTitleCase_PartitionerPrecedence(t *testing.T) {
	input := "hello_world"

	// 1. Default behavior (should split on space, but input has underscore, so one word)
	defaultWords, err := strings2.ParseTitleCase(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedDefaultStr := `[]strings2.Word{"hello_world"}`
	if fmt.Sprintf("%#v", defaultWords) != expectedDefaultStr {
		t.Errorf("default behavior: expected %s, got %#v", expectedDefaultStr, defaultWords)
	}

	// 2. Custom partitioner behavior (should override space delimiter with underscore)
	customWords, err := strings2.ParseTitleCase(input, strings2.WithSnakeCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedCustomStr := `[]strings2.Word{"hello", "world"}`
	if fmt.Sprintf("%#v", customWords) != expectedCustomStr {
		t.Errorf("custom partitioner behavior: expected %s, got %#v", expectedCustomStr, customWords)
	}
}

func TestParseTitleCaseToParts_PartitionerPrecedence(t *testing.T) {
	input := "hello_world"

	// 1. Default behavior
	defaultParts, err := strings2.ParseTitleCaseToParts(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
    if len(defaultParts) != 1 {
        t.Errorf("expected default parts len 1, got %d", len(defaultParts))
    }

	// 2. Custom partitioner behavior (WithSnakeCasePartitioner drops the delimiter underscore by default)
	customParts, err := strings2.ParseTitleCaseToParts(input, strings2.WithSnakeCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
    if len(customParts) != 2 {
        t.Errorf("expected custom parts len 2, got %d", len(customParts))
    }
}
