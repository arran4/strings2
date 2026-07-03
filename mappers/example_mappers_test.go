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
	reversed := strings2.Map(words, mappers.Reverse)
	result, _ := strings2.WordsToFormattedCase(reversed, strings2.OptionDelimiter(" "), strings2.OptionCaseMode(strings2.CMVerbatim))
	fmt.Println(result)

	// Alternatively, pass it as an option to standard formatting:
	result2, _ := strings2.ToCamel("hello world from strings2", strings2.OptionMap(mappers.Reverse))
	fmt.Println(result2)

	// Output:
	// strings2 from world hello
	// strings2FromWorldHello
}

func ExampleFilter() {
	// Filtering out words containing digits via options
	filterNumbers := mappers.Filter(func(w strings2.Word) bool {
		return !strings.ContainsAny(w.String(), "0123456789")
	})

	result, _ := strings2.ToSnake("hello 123 world", strings2.OptionMap(filterNumbers))
	fmt.Println(result)

	// Output: hello_world
}

func ExampleAcronym() {
	// Convert to acronym via options
	result, _ := strings2.ToPascal("National Aeronautics and Space Administration", strings2.OptionMap(mappers.Acronym))
	fmt.Println(result)

	// Output: NAS
}
