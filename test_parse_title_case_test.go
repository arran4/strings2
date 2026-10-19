package strings2_test

import (
	"reflect"
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
	expectedDefault := []strings2.Word{strings2.SingleCaseWord("hello_world")}
	if !reflect.DeepEqual(defaultWords, expectedDefault) {
		t.Errorf("default behavior: expected %#v, got %#v", expectedDefault, defaultWords)
	}

	// 2. Custom partitioner behavior (should override space delimiter with underscore)
	customWords, err := strings2.ParseTitleCase(input, strings2.WithSnakeCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedCustom := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.SingleCaseWord("world")}
	if !reflect.DeepEqual(customWords, expectedCustom) {
		t.Errorf("custom partitioner behavior: expected %#v, got %#v", expectedCustom, customWords)
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
	expectedDefaultWords := []strings2.Word{strings2.SingleCaseWord("hello_world")}
	if !reflect.DeepEqual(defaultWords, expectedDefaultWords) {
		t.Errorf("default parts to words: expected %#v, got %#v", expectedDefaultWords, defaultWords)
	}

	// 2. Custom partitioner behavior (WithSnakeCasePartitioner splits by underscore)
	customParts, err := strings2.ParseTitleCaseToParts(input, strings2.WithSnakeCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	customWords := strings2.PartsToWords(customParts, &strings2.ParserConfig{})
	expectedCustomWords := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.SingleCaseWord("world")}
	if !reflect.DeepEqual(customWords, expectedCustomWords) {
		t.Errorf("custom parts to words: expected %#v, got %#v", expectedCustomWords, customWords)
	}
}
