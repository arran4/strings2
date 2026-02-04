package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Empty string
	r, size := utf8.DecodeRuneInString("")
	fmt.Printf("Empty: r=%q (%d), size=%d, isError=%v\n", r, r, size, r == utf8.RuneError)

	// Invalid string
	r, size = utf8.DecodeRuneInString("\xff")
	fmt.Printf("Invalid: r=%q (%d), size=%d, isError=%v\n", r, r, size, r == utf8.RuneError)

    // Valid string
    r, size = utf8.DecodeRuneInString("a")
    fmt.Printf("Valid: r=%q (%d), size=%d, isError=%v\n", r, r, size, r == utf8.RuneError)
}
