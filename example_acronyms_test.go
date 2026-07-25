package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

func ExampleOptionAcronymList() {
	input := "Http Json Config"

	result, _ := strings2.ToSnake(input, strings2.OptionAcronymList([]string{"HTTP", "JSON"}))
	fmt.Println(result)

	// Output:
	// HTTP_JSON_Config
}
