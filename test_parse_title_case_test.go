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

	defaultWords := strings2.PartsToWords(defaultParts, &strings2.ParserConfig{})
	expectedDefaultWordsStr := `[]strings2.Word{"hello_world"}`
	if fmt.Sprintf("%#v", defaultWords) != expectedDefaultWordsStr {
		t.Errorf("default parts to words: expected %s, got %#v", expectedDefaultWordsStr, defaultWords)
	}

	// 2. Custom partitioner behavior (WithSnakeCasePartitioner splits by underscore)
	customParts, err := strings2.ParseTitleCaseToParts(input, strings2.WithSnakeCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	customWords := strings2.PartsToWords(customParts, &strings2.ParserConfig{})
	expectedCustomWordsStr := `[]strings2.Word{"hello", "world"}`
	if fmt.Sprintf("%#v", customWords) != expectedCustomWordsStr {
		t.Errorf("custom parts to words: expected %s, got %#v", expectedCustomWordsStr, customWords)
	}
}
