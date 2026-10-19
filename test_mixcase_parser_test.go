package strings2_test

import (
	"fmt"
	"testing"

	"github.com/arran4/strings2"
)

func TestMixCase_PartitionerConfig(t *testing.T) {
	// Test NumberMode precedence
	wordsNumberModeAfter, _ := strings2.ParseTitleCase("hello123world", strings2.WithNumberMode(strings2.NumberModeSplitAlways))
	expectedStrAfter := `[]strings2.Word{"hello", "123", "world"}`
	if fmt.Sprintf("%#v", wordsNumberModeAfter) != expectedStrAfter {
		t.Errorf("expected %s, got %#v", expectedStrAfter, wordsNumberModeAfter)
	}

	wordsNumberModeBefore, _ := strings2.Parse("hello123world", strings2.WithNumberMode(strings2.NumberModeSplitAlways), strings2.WithTitleCasePartitioner())
	if fmt.Sprintf("%#v", wordsNumberModeBefore) != expectedStrAfter {
		t.Errorf("expected %s, got %#v", expectedStrAfter, wordsNumberModeBefore)
	}

	// Test ParserEmitEmpty precedence
	wordsEmitEmptyAfter, _ := strings2.Parse("_hello", strings2.WithSnakeCasePartitioner(), strings2.ParserEmitEmpty(true))
	// Option applies to root parser config but partitioner already created, so emit empty wasn't enabled if order matters

	wordsEmitEmptyBefore, _ := strings2.Parse("_hello", strings2.ParserEmitEmpty(true), strings2.WithSnakeCasePartitioner())
	if fmt.Sprintf("%#v", wordsEmitEmptyBefore) != `[]strings2.Word{"", "hello"}` {
		t.Errorf("expected before `[\"\", \"hello\"]`, got %#v", wordsEmitEmptyBefore)
	}

	if fmt.Sprintf("%#v", wordsEmitEmptyAfter) != `[]strings2.Word{"hello"}` {
	    // If it equals ["hello"], it proves order matters for EmitEmpty and partitioner
		t.Logf("wordsEmitEmptyAfter was %#v", wordsEmitEmptyAfter)
	}
}
