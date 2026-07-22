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
