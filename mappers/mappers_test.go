package mappers_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func TestMapReverse(t *testing.T) {
	words := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.ExactCaseWord("World")}

	result := mappers.Reverse(words)
	expected := []strings2.Word{strings2.ExactCaseWord("World"), strings2.SingleCaseWord("hello")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}
}

func TestMapFilter(t *testing.T) {
	words := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.SeparatorWord("-"), strings2.ExactCaseWord("World"), nil}

	filterFn := mappers.Filter(func(w strings2.Word) bool {
		return !strings.Contains(w.String(), "-")
	})

	result := filterFn(words)
	expected := []strings2.Word{strings2.SingleCaseWord("hello"), strings2.ExactCaseWord("World")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}

	filterFnNil := mappers.Filter(nil)
	resultNil := filterFnNil(words)
	if !reflect.DeepEqual(words, resultNil) {
		t.Errorf("Expected %#v, got %#v", words, resultNil)
	}
}

func TestMapAcronym(t *testing.T) {
	words := []strings2.Word{
		strings2.SingleCaseWord("national"),
		strings2.SeparatorWord(" "),
		strings2.SingleCaseWord("aeronautics"),
		strings2.SeparatorWord(" "),
		strings2.SingleCaseWord("space"),
		nil,
	}

	result := mappers.Acronym(words)
	expected := []strings2.Word{strings2.AcronymWord("NAS")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}

	emptyResult := mappers.Acronym([]strings2.Word{strings2.SeparatorWord(" ")})
	if emptyResult != nil {
		t.Errorf("Expected nil, got %#v", emptyResult)
	}
}
