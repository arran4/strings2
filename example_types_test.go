package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

func ExampleUpperCaseFirst() {
	input := "hello world"
	res := strings2.UpperCaseFirst(input)
	fmt.Println(res)
	// Output: Hello world
}

func ExampleLowerCaseFirst() {
	input := "Hello World"
	res := strings2.LowerCaseFirst(input)
	fmt.Println(res)
	// Output: hello World
}

func ExampleToDarwinCase() {
	words1, _ := strings2.Parse("hello world")
	fmt.Println(strings2.Must(strings2.ToDarwinCase(words1)))
	words2, _ := strings2.Parse("camelCaseInput")
	fmt.Println(strings2.Must(strings2.ToDarwinCase(words2)))
	words3, _ := strings2.Parse("mixed-UP-Kebab")
	fmt.Println(strings2.Must(strings2.ToDarwinCase(words3)))
	// Output:
	// Hello_World
	// Camel_Case_Input
	// Mixed_Up_Kebab
}

func ExampleToTitleCase() {
	words := []strings2.Word{
		strings2.SingleCaseWord("the"),
		strings2.SingleCaseWord("lord"),
		strings2.SingleCaseWord("of"),
		strings2.SingleCaseWord("the"),
		strings2.SingleCaseWord("rings"),
	}
	result, _ := strings2.ToTitleCase(words)
	fmt.Println(result)
	// Output: The Lord of the Rings
}

func ExampleToTitleCase_screaming() {
	words := []strings2.Word{
		strings2.SingleCaseWord("A"),
		strings2.SingleCaseWord("NEW"),
		strings2.SingleCaseWord("HOPE"),
	}
	result, _ := strings2.ToTitleCase(words)
	fmt.Println(result)
	// Output: A New Hope
}
