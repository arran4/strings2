package mappers_test

import (
	"fmt"
	"strings"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func ExampleReverseWords() {
	// Reversing words via standard formatting parsing arguments
	result, _ := strings2.ToCamel("hello world from strings2", strings2.WordMapper(mappers.ReverseWords))
	fmt.Println(result)

	// Output:
	// strings2FromWorldHello
}

func ExampleFilterWords() {
	// Filtering out words containing digits natively via options
	filterNumbers := mappers.FilterWords(func(w strings2.Word) bool {
		return !strings.ContainsAny(w.String(), "0123456789")
	})

	result, _ := strings2.ToSnake("hello 123 world", strings2.WordMapper(filterNumbers))
	fmt.Println(result)

	// Output: hello_world
}

func ExampleAcronymify() {
	// Convert to acronym via options natively by mapping parts
	// This example shows how to convert a standard title-cased sentence into an acronym word by applying a SubPartMapper.
	result, _ := strings2.ToFormattedString("National Aeronautics and Space Administration", strings2.SubPartMapper(mappers.Acronymify), strings2.OptionCaseMode(strings2.CMVerbatim), strings2.OptionDelimiter(""))
	fmt.Println(result)

	// Output: NAASA
}

func ExampleAcronymify_camelCase() {
	// Acronymify operates natively on subparts and properly detects camelCase transitions
	// This example shows that Acronymify can correctly identify implicit word boundaries within camelCase strings and extract the acronym without manual tokenization.
	result, _ := strings2.ToFormattedString("nationalAeronauticsAndSpaceAdministration", strings2.SubPartMapper(mappers.Acronymify), strings2.OptionCaseMode(strings2.CMVerbatim), strings2.OptionDelimiter(""))
	fmt.Println(result)

	// Output: NAASA
}

func ExampleFromCamelToSnake_mapping() {
	// Reversing words via FromXToY formatter variants
	result, _ := strings2.FromCamelToSnake("helloWorldFromStrings2", strings2.WordMapper(mappers.ReverseWords))
	fmt.Println(result)

	// Output: Strings2_From_World_hello
}
