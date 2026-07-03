package strings2

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	words := []Word{SingleCaseWord("hello"), ExactCaseWord("World")}

	result := Map(words, nil)
	if !reflect.DeepEqual(words, result) {
		t.Errorf("Expected %#v, got %#v", words, result)
	}

	result = Map(words, func(w []Word) []Word { return MapReverse(w) })
	expected := []Word{ExactCaseWord("World"), SingleCaseWord("hello")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}
}

func TestMapFilter(t *testing.T) {
	words := []Word{SingleCaseWord("hello"), SeparatorWord("-"), ExactCaseWord("World"), nil}

	filterFn := MapFilter(func(w Word) bool {
		return !strings.Contains(w.String(), "-")
	})

	result := filterFn(words)
	expected := []Word{SingleCaseWord("hello"), ExactCaseWord("World")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}

	filterFnNil := MapFilter(nil)
	resultNil := filterFnNil(words)
	if !reflect.DeepEqual(words, resultNil) {
		t.Errorf("Expected %#v, got %#v", words, resultNil)
	}
}

func TestMapAcronym(t *testing.T) {
	words := []Word{
		SingleCaseWord("national"),
		SeparatorWord(" "),
		SingleCaseWord("aeronautics"),
		SeparatorWord(" "),
		SingleCaseWord("space"),
		nil,
	}

	result := MapAcronym(words)
	expected := []Word{AcronymWord("NAS")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}

	emptyResult := MapAcronym([]Word{SeparatorWord(" ")})
	if emptyResult != nil {
		t.Errorf("Expected nil, got %#v", emptyResult)
	}
}
