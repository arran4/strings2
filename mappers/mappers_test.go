package mappers_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func TestMapReverseWords(t *testing.T) {
	words := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.ExactCaseWord("World")}

	result := mappers.ReverseWords(words)
	expected := []strings2.Word{strings2.ExactCaseWord("World"), strings2.SingleCaseWord("hello")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}
}

func TestMapFilterWords(t *testing.T) {
	words := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.SeparatorWord("-"), strings2.ExactCaseWord("World"), nil}

	filterFn := mappers.FilterWords(func(w strings2.Word) bool {
		return !strings.Contains(w.String(), "-")
	})

	result := filterFn(words)
	expected := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.ExactCaseWord("World")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}

	filterFnNil := mappers.FilterWords(nil)
	resultNil := filterFnNil(words)
	if !reflect.DeepEqual(words, resultNil) {
		t.Errorf("Expected %#v, got %#v", words, resultNil)
	}
}

func TestMapAcronymify(t *testing.T) {
	subs := Must(strings2.StringToSubParts("national aeronautics space"))
	result := mappers.Acronymify(subs)

	if len(result) != 3 {
		t.Fatalf("Expected 3 subparts, got %d", len(result))
	}
	if result[0].Rune() != 'N' || result[1].Rune() != 'A' || result[2].Rune() != 'S' {
		t.Errorf("Expected NAS, got %c%c%c", result[0].Rune(), result[1].Rune(), result[2].Rune())
	}

	emptySubs := Must(strings2.StringToSubParts(" "))
	emptyResult := mappers.Acronymify(emptySubs)
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty result, got %#v", emptyResult)
	}

	// Camel case test
	camelSubs := Must(strings2.StringToSubParts("nationalAeronauticsSpace"))
	camelResult := mappers.Acronymify(camelSubs)
	if len(camelResult) != 3 {
		t.Fatalf("Expected 3 subparts, got %d", len(camelResult))
	}
	if camelResult[0].Rune() != 'N' || camelResult[1].Rune() != 'A' || camelResult[2].Rune() != 'S' {
		t.Errorf("Expected NAS, got %c%c%c", camelResult[0].Rune(), camelResult[1].Rune(), camelResult[2].Rune())
	}
}

// Compile-time check to ensure our mappers conform to the interface
var _ strings2.WordMapper = mappers.ReverseWords
var _ strings2.SubPartMapper = mappers.Acronymify
var _ strings2.PartMapper = mappers.ReverseParts
// Filter returns a WordMapper
var _ strings2.WordMapper = mappers.FilterWords(func(w strings2.Word) bool { return true })
