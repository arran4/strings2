package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

func ExampleToCamel() {
	input := "hello_world"
	res, _ := strings2.ToCamel(input)
	fmt.Println(res)
	// Output: helloWorld
}

func ExampleToSnake() {
	input := "helloWorld"
	// By default Snake uses Verbatim casing, so "World" retains its uppercase 'W'.
	res, _ := strings2.ToSnake(input)
	fmt.Println(res)
	// Output: hello_World
}

func ExampleFromCamelToSnake() {
	input := "helloWorld"
	// By default Snake uses Verbatim casing.
	res, _ := strings2.FromCamelToSnake(input)
	fmt.Println(res)
	// Output: hello_World
}

func ExampleFromCamelToKebab() {
	input := "helloWorld"
	// By default Kebab uses Verbatim casing.
	res, _ := strings2.FromCamelToKebab(input)
	fmt.Println(res)
	// Output: hello-World
}

func ExampleToLowerCamel() {
	res, _ := strings2.ToLowerCamel("some_kind_of_example")
	fmt.Println(res)
	// Output:
	// someKindOfExample
}

func ExampleToScreamingSnake() {
	res, _ := strings2.ToScreamingSnake("some_kind_of_example")
	fmt.Println(res)
	// Output:
	// SOME_KIND_OF_EXAMPLE
}

func ExampleToScreamingKebab() {
	res, _ := strings2.ToScreamingKebab("some_kind_of_example")
	fmt.Println(res)
	// Output:
	// SOME-KIND-OF-EXAMPLE
}

func ExampleToDelimited() {
	res, _ := strings2.ToDelimited("some_kind_of_example", '.')
	fmt.Println(res)
	// Output:
	// some.kind.of.example
}

func ExampleToScreamingDelimited() {
	res, _ := strings2.ToScreamingDelimited("some_kind.of_example", '.', "", true)
	fmt.Println(res)
	// Output:
	// SOME.KIND.OF.EXAMPLE
}

func ExampleToDelimited_whispering() {
	res, _ := strings2.ToDelimited("someKind", '.')
	fmt.Println(res)
	// Output:
	// some.kind
}

func ExampleToScreamingDelimited_whispering() {
	res, _ := strings2.ToScreamingDelimited("someKind", '.', "", false)
	fmt.Println(res)
	// Output:
	// some.kind
}

func ExampleOptionMixCaseSupport() {
	// FirstUpperCaseWord normalizes case to lowercase before applying First Upper.
	// Therefore, OptionMixCaseSupport is ineffective on it.
	input := []strings2.Word{strings2.FirstUpperCaseWord("camelCase")}
	res := strings2.ToFormattedCase(input, strings2.OptionMixCaseSupport(), strings2.OptionDelimiter("-"))
	fmt.Println(res)

	// ExactCaseWord does not normalize case, so MixCaseSupport works perfectly.
	input2 := []strings2.Word{strings2.ExactCaseWord("camelCase")}
	res2 := strings2.ToFormattedCase(input2, strings2.OptionMixCaseSupport(), strings2.OptionDelimiter("-"))
	fmt.Println(res2)

	// Output:
	// Camelcase
	// camel-Case
}
