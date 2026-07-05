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

// stringToPart is a test helper since Part building is internal/verbose
func stringToPart(s string, isSeparator bool) strings2.Part {
	var subs []strings2.SubPart
	for _, r := range s {
		b := strings2.BaseSubPart{Val: r}
		if isSeparator {
			subs = append(subs, strings2.SpaceSubPart{BaseSubPart: b})
		} else {
			subs = append(subs, strings2.LetterSubPart{BaseSubPart: b})
		}
	}
	if isSeparator {
		return &strings2.SeparatorPart{BasePart: strings2.BasePart{Subs: subs}}
	}
	return &strings2.WordPart{BasePart: strings2.BasePart{Subs: subs}}
}

func TestMapAcronymify(t *testing.T) {
	parts := []strings2.Part{
		stringToPart("national", false),
		stringToPart(" ", true),
		stringToPart("aeronautics", false),
		stringToPart(" ", true),
		stringToPart("space", false),
		nil,
	}

	result := mappers.Acronymify(parts)
	if len(result) != 1 || result[0].String() != "NAS" {
		t.Errorf("Expected NAS, got %#v", result)
	}

	emptyResult := mappers.Acronymify([]strings2.Part{stringToPart(" ", true)})
	if emptyResult != nil {
		t.Errorf("Expected nil, got %#v", emptyResult)
	}
}

// Compile-time check to ensure our mappers conform to the interface
var _ strings2.WordMapper = mappers.ReverseWords
var _ strings2.PartMapper = mappers.Acronymify
var _ strings2.PartMapper = mappers.ReverseParts
// Filter returns a WordMapper
var _ strings2.WordMapper = mappers.FilterWords(func(w strings2.Word) bool { return true })
