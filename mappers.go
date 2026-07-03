package strings2

// WordMapper applies a mapping function to a slice of Words.
type WordMapper func([]Word) []Word

// OptionMap creates an Option that applies the given mapping functions.
func OptionMap(mappers ...WordMapper) Option {
	return func(cfg *caseConfig) {
		cfg.mappers = append(cfg.mappers, mappers...)
	}
}

// Map applies a series of mapping functions to a slice of Words.
// It allows modification of the contents such as reversing, filtering, or custom transformations.
func Map(words []Word, mappers ...WordMapper) []Word {
	for _, m := range mappers {
		if m != nil {
			words = m(words)
		}
	}
	return words
}
