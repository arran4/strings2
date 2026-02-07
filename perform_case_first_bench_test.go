package strings2

import (
	"testing"
	"unicode"
)

func BenchmarkPerformCaseFirst(b *testing.B) {
	s := "test"
	fn := unicode.ToUpper
	for i := 0; i < b.N; i++ {
		performCaseFirst(s, fn)
	}
}

func BenchmarkPerformCaseFirst_Long(b *testing.B) {
	s := "teststringwithmorecharacters"
	fn := unicode.ToUpper
	for i := 0; i < b.N; i++ {
		performCaseFirst(s, fn)
	}
}
