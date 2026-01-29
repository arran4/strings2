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
