package strings2

import (
	"testing"
)

var benchWords = []Word{
	SingleCaseWord("hello"),
	SingleCaseWord("world"),
	SingleCaseWord("foo"),
	SingleCaseWord("bar"),
}

func BenchmarkToKebabCase(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ToKebabCase(benchWords)
	}
}

func BenchmarkToKebabCase_WithOptions(b *testing.B) {
	b.ReportAllocs()
	opts := []Option{OptionCaseMode(CMScreaming), OptionUpperIndicator("--")}
	for i := 0; i < b.N; i++ {
		ToKebabCase(benchWords, opts...)
	}
}

func BenchmarkToKebabCase_WithManyOptions(b *testing.B) {
	b.ReportAllocs()
	opts := []Option{
		OptionCaseMode(CMScreaming),
		OptionUpperIndicator("--"),
		OptionFirstUpper(),
		OptionFirstLower(),
	}
	for i := 0; i < b.N; i++ {
		ToKebabCase(benchWords, opts...)
	}
}
