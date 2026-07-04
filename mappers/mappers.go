package mappers

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/arran4/strings2"
)

// ReverseWords is a mapping function that reverses the order of the words.
func ReverseWords(words []strings2.Word) []strings2.Word {
	reversed := make([]strings2.Word, len(words))
	for i, w := range words {
		reversed[len(words)-1-i] = w
	}
	return reversed
}

// FilterWords returns a mapping function that keeps only the words for which the keep function returns true.
func FilterWords(keep func(strings2.Word) bool) func([]strings2.Word) []strings2.Word {
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

// Acronymify creates an acronym by taking the first letter of each part, converting it to an UpperCaseWord
func Acronymify(parts []strings2.Part) []strings2.Part {
	var b strings.Builder
	for _, p := range parts {
		if p == nil {
			continue
		}
		if _, ok := p.(*strings2.SeparatorPart); ok {
			continue
		}
		s := p.String()
		if len(s) > 0 {
			r, _ := utf8.DecodeRuneInString(s)
			b.WriteRune(unicode.ToUpper(r))
		}
	}
	if b.Len() > 0 {
		// Construct SubParts manually since BasePart operates on SubParts
		var subs []strings2.SubPart
		for _, r := range b.String() {
			subs = append(subs, strings2.LetterSubPart{BaseSubPart: strings2.BaseSubPart{Val: r}})
		}
		return []strings2.Part{&strings2.WordPart{BasePart: strings2.BasePart{Subs: subs}}}
	}
	return nil
}

// ReverseParts is a mapping function that reverses the order of the parts.
func ReverseParts(parts []strings2.Part) []strings2.Part {
	reversed := make([]strings2.Part, len(parts))
	for i, p := range parts {
		reversed[len(parts)-1-i] = p
	}
	return reversed
}
