package strings2

// WordMapper applies a mapping function to a slice of Words.
type WordMapper func([]Word) []Word

// ApplyWordMappers applies a series of WordMapper functions to a slice of Words.
func ApplyWordMappers(words []Word, mappers ...WordMapper) []Word {
	for _, m := range mappers {
		if m != nil {
			words = m(words)
		}
	}
	return words
}

// PartMapper applies a mapping function to a slice of Parts.
type PartMapper func([]Part) []Part

// ApplyPartMappers applies a series of PartMapper functions to a slice of Parts.
func ApplyPartMappers(parts []Part, mappers ...PartMapper) []Part {
	for _, m := range mappers {
		if m != nil {
			parts = m(parts)
		}
	}
	return parts
}

// SubPartMapper applies a mapping function to a slice of SubParts.
type SubPartMapper func([]SubPart) []SubPart

// ApplySubPartMappers applies a series of SubPartMapper functions to a slice of SubParts.
func ApplySubPartMappers(subs []SubPart, mappers ...SubPartMapper) []SubPart {
	for _, m := range mappers {
		if m != nil {
			subs = m(subs)
		}
	}
	return subs
}
