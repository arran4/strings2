package strings2_test

import (
	"fmt"
	"strings"

	"github.com/arran4/strings2"
)

func ExampleMap_reverse() {
	words, _ := strings2.Parse("hello world from strings2")

	// Reverse the words
	reversed := strings2.Map(words, strings2.MapReverse)

	result, _ := strings2.WordsToFormattedCase(reversed, strings2.OptionDelimiter(" "), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)
	// Output: strings2 from world hello
}

func ExampleMap_filter() {
	words, _ := strings2.Parse("hello 123 world")

	// Filter out words containing digits
	noDigits := strings2.Map(words, strings2.MapFilter(func(w strings2.Word) bool {
		return !strings.ContainsAny(w.String(), "0123456789")
	}))

	result, _ := strings2.WordsToFormattedCase(noDigits, strings2.OptionDelimiter(" "), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)
	// Output: hello world
}

func ExampleMap_acronym() {
	words, _ := strings2.Parse("National Aeronautics and Space Administration")

	// Convert to acronym
	acronymWords := strings2.Map(words, strings2.MapAcronym)

	result, _ := strings2.WordsToFormattedCase(acronymWords, strings2.OptionDelimiter(""), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)
	// Output: NAASA
}
