package strings2_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arran4/strings2"
)

func ExampleOptionAcronymList() {
	input := "Http Json Config"

	result, _ := strings2.ToSnake(input, strings2.OptionAcronymList([]string{"HTTP", "JSON"}))
	fmt.Println(result)

	// Output:
	// HTTP_JSON_Config
}

func ExampleLoadAcronymsFromFile() {
	// Create a temporary file containing acronyms
	dir, _ := os.MkdirTemp("", "acronyms")
	defer os.RemoveAll(dir)
	file := filepath.Join(dir, "acronyms.txt")
	os.WriteFile(file, []byte("HTTP\nJSON\n"), 0644)

	mapper, _ := strings2.LoadAcronymsFromFile(file)

	input := "Http Json Config"
	result, _ := strings2.ToSnake(input, mapper)
	fmt.Println(result)

	// Output:
	// HTTP_JSON_Config
}
