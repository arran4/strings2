package mappers

import (
	"unicode"

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

// Acronymify creates an acronym by taking the first letter of each word and discarding the rest.
//
// Analysis of previous implementation complexity:
// The original implementation as a PartMapper was overly complex because it forced manual
// reconstruction of the internal parsing AST (SubParts and Parts). It extracted letters into a
// strings.Builder and then manually re-lexed them into a new WordPart.
// By implementing it as a SubPartMapper, it acts purely as an early-stage filter.
// It filters the stream to keep only the capitalized first letter of each word boundary.
// The downstream pipeline (Partitioner and Word Classifier) will then naturally group these
// adjacent letters into a single UpperCaseWord or AcronymWord, completely eliminating the need
// for manual AST reconstruction.
func Acronymify(subs []strings2.SubPart) []strings2.SubPart {
	var result []strings2.SubPart
	inWord := false
	for i, sp := range subs {
		if sp == nil {
			inWord = false
			continue
		}
		if sp.IsLetter() {
			isNewWord := !inWord
			if inWord && i > 0 {
				prev := subs[i-1]
				// Camel case boundary: lower to Upper
				if prev != nil && prev.IsLower() && sp.IsUpper() {
					isNewWord = true
				}
				// Camel case boundary: Upper to Upper to lower (e.g., PDFLoader -> P, L)
				if prev != nil && prev.IsUpper() && sp.IsUpper() && i+1 < len(subs) && subs[i+1] != nil && subs[i+1].IsLower() {
					isNewWord = true
				}
			}
			if isNewWord {
				// Convert the first letter to uppercase
				if sp.IsUpper() {
					result = append(result, sp)
				} else {
					result = append(result, strings2.LetterSubPart{
						BaseSubPart: strings2.BaseSubPart{Val: unicode.ToUpper(sp.Rune())},
					})
				}
				inWord = true
			}
		} else {
			inWord = false
		}
	}
	return result
}

// ReverseParts is a mapping function that reverses the order of the parts.
func ReverseParts(parts []strings2.Part) []strings2.Part {
	reversed := make([]strings2.Part, len(parts))
	for i, p := range parts {
		reversed[len(parts)-1-i] = p
	}
	return reversed
}
