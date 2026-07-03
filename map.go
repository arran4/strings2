package strings2

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Map applies a series of mapping functions to a slice of Words.
// It allows modification of the contents such as reversing, filtering, or custom transformations.
func Map(words []Word, mappers ...func([]Word) []Word) []Word {
	for _, m := range mappers {
		if m != nil {
			words = m(words)
		}
	}
	return words
}

// MapReverse is a mapping function that reverses the order of the words.
func MapReverse(words []Word) []Word {
	reversed := make([]Word, len(words))
	for i, w := range words {
		reversed[len(words)-1-i] = w
	}
	return reversed
}

// MapFilter returns a mapping function that keeps only the words for which the keep function returns true.
func MapFilter(keep func(Word) bool) func([]Word) []Word {
	if keep == nil {
		return func(words []Word) []Word { return words }
	}
	return func(words []Word) []Word {
		var filtered []Word
		for _, w := range words {
			if w != nil && keep(w) {
				filtered = append(filtered, w)
			}
		}
		return filtered
	}
}

// MapAcronym is a mapping function that creates an acronym by taking the first letter of each word
// and combining them into a single AcronymWord.
func MapAcronym(words []Word) []Word {
	var b strings.Builder
	for _, w := range words {
		if w == nil {
			continue
		}
		// Ignore separators
		if _, ok := w.(SeparatorWord); ok {
			continue
		}
		s := w.String()
		if len(s) > 0 {
			r, _ := utf8.DecodeRuneInString(s)
			b.WriteRune(unicode.ToUpper(r))
		}
	}
	if b.Len() > 0 {
		return []Word{AcronymWord(b.String())}
	}
	return nil
}
