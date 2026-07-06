package strings2_test

import (
	"fmt"

	"github.com/arran4/strings2"
)

func ExampleDelimiterDetector() {
	input := "foo::bar--baz"

	// Create a partitioner that looks for specific multi-rune sequences ("::" and "--")
	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		DelimiterDetector: func(subs []strings2.SubPart, index int) int {
			// Look for double colon "::"
			if index+1 < len(subs) && subs[index].Rune() == ':' && subs[index+1].Rune() == ':' {
				return 2
			}
			// Look for double dash "--"
			if index+1 < len(subs) && subs[index].Rune() == '-' && subs[index+1].Rune() == '-' {
				return 2
			}
			return 0
		},
		PreserveSep: true,
	})

	words, _ := strings2.Parse(input, partitioner)
	for i, w := range words {
		fmt.Printf("Word %d: %T %q\n", i, w, w.String())
	}

	// Output:
	// Word 0: strings2.SingleCaseWord "foo"
	// Word 1: strings2.SeparatorWord "::"
	// Word 2: strings2.SingleCaseWord "bar"
	// Word 3: strings2.SeparatorWord "--"
	// Word 4: strings2.SingleCaseWord "baz"
}
