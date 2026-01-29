package strings2

import (
	"testing"
)

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
