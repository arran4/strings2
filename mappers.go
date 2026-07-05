package strings2

// WordMapper applies a mapping function to a slice of Words.
type WordMapper func([]Word) []Word

// PartMapper applies a mapping function to a slice of Parts.
type PartMapper func([]Part) []Part

// SubPartMapper applies a mapping function to a slice of SubParts.
type SubPartMapper func([]SubPart) []SubPart
// trigger commit
