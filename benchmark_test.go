package strings2

import (
	"testing"
)

func BenchmarkToFormattedCase(b *testing.B) {
	words := []Word{
		SingleCaseWord("one"),
		SingleCaseWord("two"),
		SingleCaseWord("three"),
		SingleCaseWord("four"),
		SingleCaseWord("five"),
		SingleCaseWord("six"),
		SingleCaseWord("seven"),
		SingleCaseWord("eight"),
		SingleCaseWord("nine"),
		SingleCaseWord("ten"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToFormattedCase(words, OptionDelimiter("-"))
	}
}
