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
