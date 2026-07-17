package mappers

import (
	"reflect"
	"testing"

	"github.com/arran4/strings2"
)

func TestMapAcronyms(t *testing.T) {
	words := []strings2.Word{
		strings2.SingleCaseWord("user"),
		strings2.SingleCaseWord("id"),
	}

	mapper := MapAcronyms("ID")
	result := mapper(words)

	expected := []strings2.Word{
		strings2.SingleCaseWord("user"),
		strings2.AcronymWord("ID"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
