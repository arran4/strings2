package mappers_test

import (
	"fmt"
	"strings"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func ExampleReverse() {
	words, _ := strings2.Parse("hello world from strings2")

	// Reverse the words
	reversed := strings2.Map(words, mappers.Reverse)

	result, _ := strings2.WordsToFormattedCase(reversed, strings2.OptionDelimiter(" "), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)
	// Output: strings2 from world hello
}

func ExampleFilter() {
	words, _ := strings2.Parse("hello 123 world")

	// Filter out words containing digits
	noDigits := strings2.Map(words, mappers.Filter(func(w strings2.Word) bool {
		return !strings.ContainsAny(w.String(), "0123456789")
	}))

	result, _ := strings2.WordsToFormattedCase(noDigits, strings2.OptionDelimiter(" "), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)
	// Output: hello world
}

func ExampleAcronym() {
	words, _ := strings2.Parse("National Aeronautics and Space Administration")

	// Convert to acronym
	acronymWords := strings2.Map(words, mappers.Acronym)

	result, _ := strings2.WordsToFormattedCase(acronymWords, strings2.OptionDelimiter(""), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)
	// Output: NAASA
}
