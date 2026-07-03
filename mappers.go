package strings2

// Mapper applies a series of mapping functions to a slice of Words.
type Mapper func([]Word) []Word

// OptionMap creates an Option that applies the given mapping functions.
func OptionMap(mappers ...Mapper) Option {
	return func(cfg *caseConfig) {
		cfg.mappers = append(cfg.mappers, mappers...)
	}
}

// Map applies a series of mapping functions to a slice of Words.
// It allows modification of the contents such as reversing, filtering, or custom transformations.
func Map(words []Word, mappers ...Mapper) []Word {
	for _, m := range mappers {
		if m != nil {
			words = m(words)
		}
	}
	return words
}
