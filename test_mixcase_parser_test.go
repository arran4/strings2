package strings2_test

import (
	"reflect"
	"testing"

	"github.com/arran4/strings2"
)

func TestMixCase_PartitionerConfig(t *testing.T) {
	// Test NumberMode precedence

	// Case 1: Partitioner before NumberMode
	wordsNumberModeAfter, err := strings2.Parse("hello123world", strings2.WithTitleCasePartitioner(), strings2.WithNumberMode(strings2.NumberModeSplitAlways))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// The number mode is not passed to the TitleCasePartitioner because it was instantiated BEFORE WithNumberMode was encountered.
	// So we expect "hello123world" as a single block for default number mode behavior.
	expectedStrAfter := []strings2.Word{strings2.ExactCaseWord("hello123world")}
	if !reflect.DeepEqual(wordsNumberModeAfter, expectedStrAfter) {
		t.Errorf("expected %#v, got %#v", expectedStrAfter, wordsNumberModeAfter)
	}

	// Case 2: NumberMode before Partitioner
	wordsNumberModeBefore, err := strings2.Parse("hello123world", strings2.WithNumberMode(strings2.NumberModeSplitAlways), strings2.WithTitleCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedStrBefore := []strings2.Word{strings2.ExactCaseWord("hello"), strings2.ExactCaseWord("123"), strings2.ExactCaseWord("world")}
	if !reflect.DeepEqual(wordsNumberModeBefore, expectedStrBefore) {
		t.Errorf("expected %#v, got %#v", expectedStrBefore, wordsNumberModeBefore)
	}

	// Test ParserEmitEmpty precedence
	// Case 1: Partitioner before EmitEmpty
	wordsEmitEmptyAfter, err := strings2.Parse("_hello", strings2.WithSnakeCasePartitioner(), strings2.ParserEmitEmpty(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Option applies to root parser config but partitioner already created, so emit empty wasn't enabled for the partitioner logic.
	expectedEmitEmptyAfter := []strings2.Word{strings2.ExactCaseWord("hello")}
	if !reflect.DeepEqual(wordsEmitEmptyAfter, expectedEmitEmptyAfter) {
		t.Errorf("EmitEmpty After Partitioner: expected %#v, got %#v", expectedEmitEmptyAfter, wordsEmitEmptyAfter)
	}

	// Case 2: EmitEmpty before Partitioner
	wordsEmitEmptyBefore, err := strings2.Parse("_hello", strings2.ParserEmitEmpty(true), strings2.WithSnakeCasePartitioner())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedEmitEmptyBefore := []strings2.Word{strings2.ExactCaseWord(""), strings2.ExactCaseWord("hello")}
	if !reflect.DeepEqual(wordsEmitEmptyBefore, expectedEmitEmptyBefore) {
		t.Errorf("EmitEmpty Before Partitioner: expected %#v, got %#v", expectedEmitEmptyBefore, wordsEmitEmptyBefore)
	}
}
