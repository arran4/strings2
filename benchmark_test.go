package strings2

import (
	"testing"
)

func BenchmarkUpperCaseFirst_ASCII_Short(b *testing.B) {
	s := "test"
	for i := 0; i < b.N; i++ {
		UpperCaseFirst(s)
	}
}

func BenchmarkUpperCaseFirst_ASCII_Long(b *testing.B) {
	s := "teststringwithmorecharacters"
	for i := 0; i < b.N; i++ {
		UpperCaseFirst(s)
	}
}

func BenchmarkUpperCaseFirst_Unicode(b *testing.B) {
	s := "äpfel"
	for i := 0; i < b.N; i++ {
		UpperCaseFirst(s)
	}
}

func BenchmarkUpperCaseFirst_Empty(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		UpperCaseFirst(s)
	}
}

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

func BenchmarkToFormattedCase_Screaming(b *testing.B) {
	words := []Word{
		SingleCaseWord("hello"),
		SingleCaseWord("world"),
		SingleCaseWord("this"),
		SingleCaseWord("is"),
		SingleCaseWord("a"),
		SingleCaseWord("test"),
	}
	opts := []Option{OptionCaseMode(CMScreaming)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToFormattedCase(words, opts...)
	}
}

func BenchmarkToFormattedCase_Screaming_Mixed(b *testing.B) {
	words := []Word{
		SingleCaseWord("HeLLo"),
		SingleCaseWord("WoRLd"),
		SingleCaseWord("ThIs"),
		SingleCaseWord("IS"),
		SingleCaseWord("A"),
		SingleCaseWord("TeSt"),
	}
	opts := []Option{OptionCaseMode(CMScreaming)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToFormattedCase(words, opts...)
	}
}

func BenchmarkToFormattedCase_Whispering(b *testing.B) {
	words := []Word{
		SingleCaseWord("HELLO"),
		SingleCaseWord("WORLD"),
		SingleCaseWord("THIS"),
		SingleCaseWord("IS"),
		SingleCaseWord("A"),
		SingleCaseWord("TEST"),
	}
	opts := []Option{OptionCaseMode(CMWhispering)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToFormattedCase(words, opts...)
	}
}

func BenchmarkToFormattedCase_AllTitle(b *testing.B) {
	words := []Word{
		SingleCaseWord("hello"),
		SingleCaseWord("world"),
		SingleCaseWord("this"),
		SingleCaseWord("is"),
		SingleCaseWord("a"),
		SingleCaseWord("test"),
	}
	opts := []Option{OptionCaseMode(CMAllTitle)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToFormattedCase(words, opts...)
	}
}

func BenchmarkToFormattedCase_Default(b *testing.B) {
	words := []Word{
		SingleCaseWord("HeLLo"),
		SingleCaseWord("WoRLd"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToFormattedCase(words)
	}
}
