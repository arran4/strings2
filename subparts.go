package strings2

import (
	"unicode"
)

// SubPart represents the smallest unit of parsing, typically a single character with its properties.
type SubPart interface {
	Rune() rune
	IsDigit() bool
	IsLetter() bool
	IsUpper() bool
	IsLower() bool
	IsSpace() bool
	IsSymbol() bool
}

type BaseSubPart struct {
	Val rune
}

func (b BaseSubPart) Rune() rune     { return b.Val }
func (b BaseSubPart) IsDigit() bool  { return unicode.IsDigit(b.Val) }
func (b BaseSubPart) IsLetter() bool { return unicode.IsLetter(b.Val) }
func (b BaseSubPart) IsUpper() bool  { return unicode.IsUpper(b.Val) }
func (b BaseSubPart) IsLower() bool  { return unicode.IsLower(b.Val) }
func (b BaseSubPart) IsSpace() bool  { return unicode.IsSpace(b.Val) }
func (b BaseSubPart) IsSymbol() bool { return !b.IsDigit() && !b.IsLetter() && !b.IsSpace() }

type LetterSubPart struct{ BaseSubPart }
type DigitSubPart struct{ BaseSubPart }
type SpaceSubPart struct{ BaseSubPart }
type SymbolSubPart struct{ BaseSubPart }

// Stats contains statistics about the scanned string.
type Stats struct {
	TotalLen int
	Letters  int
	Digits   int
	Spaces   int
	Symbols  int
	Upper    int
	Lower    int

	// Histogram of specific symbols for delimiter detection
	SymbolCounts map[rune]int
}

// StringToSubParts converts a string into a slice of SubParts and generates stats.
func StringToSubParts(s string) ([]SubPart, Stats) {
	var subs []SubPart
	stats := Stats{
		SymbolCounts: make(map[rune]int),
	}

	for _, r := range s {
		stats.TotalLen++
		var sp SubPart
		b := BaseSubPart{Val: r}

		if unicode.IsLetter(r) {
			sp = LetterSubPart{b}
			stats.Letters++
			if unicode.IsUpper(r) {
				stats.Upper++
			} else if unicode.IsLower(r) {
				stats.Lower++
			}
		} else if unicode.IsDigit(r) {
			sp = DigitSubPart{b}
			stats.Digits++
		} else if unicode.IsSpace(r) {
			sp = SpaceSubPart{b}
			stats.Spaces++
		} else {
			sp = SymbolSubPart{b}
			stats.Symbols++
			stats.SymbolCounts[r]++
		}

		subs = append(subs, sp)
	}

	return subs, stats
}
