package strings2

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	words := []Word{SingleCaseWord("hello"), ExactCaseWord("World")}

	result := Map(words, nil)
	if !reflect.DeepEqual(words, result) {
		t.Errorf("Expected %#v, got %#v", words, result)
	}

	result = Map(words, func(w []Word) []Word {
		// simple reverse just for Map testing
		reversed := make([]Word, len(w))
		for i, x := range w {
			reversed[len(w)-1-i] = x
		}
		return reversed
	})
	expected := []Word{ExactCaseWord("World"), SingleCaseWord("hello")}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}
}
