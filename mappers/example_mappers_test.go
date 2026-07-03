package mappers_test

import (
	"fmt"
	"strings"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func ExampleReverse() {
	// Reversing using the mapper directly on words
	words, _ := strings2.Parse("hello world from strings2")
	// Since ApplyWordMappers is internal, directly call the mapper if applying manually
	reversed := mappers.Reverse(words)
	result, _ := strings2.WordsToFormattedCase(reversed, strings2.OptionDelimiter(" "), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)

	// Alternatively, pass it as an option to standard formatting natively:
	// NOTE: mappers apply *after* parsing, so `ToCamel` converts input to words,
	//       then applies `Reverse`, then applies Camel formatting (first lower, then title).
	result2, _ := strings2.ToCamel("hello world from strings2", strings2.WordMapper(mappers.Reverse))
	fmt.Println(result2)

	// Output:
	// strings2 from world hello
	// strings2FromWorldHello
}

func ExampleFilter() {
	// Filtering out words containing digits natively via options
	filterNumbers := mappers.Filter(func(w strings2.Word) bool {
		return !strings.ContainsAny(w.String(), "0123456789")
	})

	result, _ := strings2.ToSnake("hello 123 world", strings2.WordMapper(filterNumbers))
	fmt.Println(result)

	// Output: hello_world
}

func ExampleAcronym() {
	// Convert to acronym via options natively using Verbatim to preserve acronym casing
	result, _ := strings2.ToFormattedString("National Aeronautics and Space Administration", strings2.WordMapper(mappers.Acronym), strings2.OptionCaseMode(strings2.CMVerbatim), strings2.OptionDelimiter(""))
	fmt.Println(result)

	// Output: NAASA
}
