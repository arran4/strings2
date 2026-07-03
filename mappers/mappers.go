package mappers

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/arran4/strings2"
)

// Reverse is a mapping function that reverses the order of the words.
func Reverse(words []strings2.Word) []strings2.Word {
	reversed := make([]strings2.Word, len(words))
	for i, w := range words {
		reversed[len(words)-1-i] = w
	}
	return reversed
}

// Filter returns a mapping function that keeps only the words for which the keep function returns true.
func Filter(keep func(strings2.Word) bool) func([]strings2.Word) []strings2.Word {
	if keep == nil {
		return func(words []strings2.Word) []strings2.Word { return words }
	}
	return func(words []strings2.Word) []strings2.Word {
		var filtered []strings2.Word
		for _, w := range words {
			if w != nil && keep(w) {
				filtered = append(filtered, w)
			}
		}
		return filtered
	}
}

// Acronym is a mapping function that creates an acronym by taking the first letter of each word
// and combining them into a single AcronymWord.
func Acronym(words []strings2.Word) []strings2.Word {
	var b strings.Builder
	for _, w := range words {
		if w == nil {
			continue
		}
		// Ignore separators
		if _, ok := w.(strings2.SeparatorWord); ok {
			continue
		}
		s := w.String()
		if len(s) > 0 {
			r, _ := utf8.DecodeRuneInString(s)
			b.WriteRune(unicode.ToUpper(r))
		}
	}
	if b.Len() > 0 {
		return []strings2.Word{strings2.AcronymWord(b.String())}
	}
	return nil
}
