package strings2

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
