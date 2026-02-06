package strings2

import (
	"strings"
)

// Part represents a grouped sequence of SubParts.
type Part interface {
	String() string
	SubParts() []SubPart
}

type BasePart struct {
	Subs []SubPart
}

func (p BasePart) String() string {
	var sb strings.Builder
	for _, s := range p.Subs {
		sb.WriteRune(s.Rune())
	}
	return sb.String()
}

func (p BasePart) SubParts() []SubPart { return p.Subs }

// Concrete Part types
type WordPart struct{ BasePart }
type SeparatorPart struct{ BasePart }

// Partitioner defines a function that groups SubParts into Parts.
type Partitioner func([]SubPart) []Part

// SubPartsToParts converts SubParts to Parts using the provided Partitioner.
func SubPartsToParts(subs []SubPart, partitioner Partitioner) []Part {
	if partitioner == nil {
		return []Part{&WordPart{BasePart{Subs: subs}}}
	}
	return partitioner(subs)
}

// Common Partitioners

// SnakeCasePartitioner splits on underscore '_'.
func SnakeCasePartitioner(subs []SubPart) []Part {
	return SplitByDelimiter(subs, '_')
}

// KebabCasePartitioner splits on hyphen '-'.
func KebabCasePartitioner(subs []SubPart) []Part {
	return SplitByDelimiter(subs, '-')
}

// SplitByDelimiter is a helper to split SubParts by a specific rune delimiter.
func SplitByDelimiter(subs []SubPart, delim rune) []Part {
	return NewPartitioner(map[rune]bool{delim: true}, false, false)(subs)
}

// CamelCasePartitioner splits on case transitions.
func CamelCasePartitioner(subs []SubPart) []Part {
	return NewPartitioner(nil, true, false)(subs)
}

// NewPartitioner creates a partitioner with specific delimiters and camel case splitting enabled.
func NewPartitioner(delimiters map[rune]bool, splitCamel bool, splitNumber bool) Partitioner {
	return func(subs []SubPart) []Part {
		var parts []Part
		var current []SubPart

		for i, s := range subs {
			// Check if current rune is a delimiter
			if delimiters != nil && delimiters[s.Rune()] {
				if len(current) > 0 {
					parts = append(parts, &WordPart{BasePart{Subs: current}})
					current = nil
				}
				// We discard delimiters
				continue
			}

			// Transition check for CamelCase
			isSplit := false
			if (splitCamel || splitNumber) && i > 0 && len(current) > 0 {
				prev := subs[i-1]
				// If prev was a delimiter, we effectively started a new word, so no split check needed based on transition from it?
				// But we skipped it. So current[len-1] is the last NON-delimiter.
				// However, `CamelCasePartitioner` logic uses `subs[i-1]`.
				// If `subs[i-1]` was a delimiter, `len(current)` is 0 (handled above) OR we appended `s`.
				// Wait, if `subs[i-1]` was delimiter, we did `continue`.
				// So `current` is empty.
				// If `current` is NOT empty, then `subs[i-1]` was NOT a delimiter (or we are in a run of characters).
				// So we can safely use `prev`.

				if splitCamel {
					// lower -> Upper
					if prev.IsLower() && s.IsUpper() {
						isSplit = true
					}

					// Upper -> Upper -> lower (PDFLoader split at L)
					if i+1 < len(subs) {
						next := subs[i+1]
						if prev.IsUpper() && s.IsUpper() && next.IsLower() {
							isSplit = true
						}
					}
				}

				if splitNumber {
					// Letter -> Digit -> Split.
					// Digit -> Letter -> Split.
					if prev.IsLetter() && s.IsDigit() {
						isSplit = true
					}
					if prev.IsDigit() && s.IsLetter() {
						isSplit = true
					}
				}
			}

			if isSplit {
				if len(current) > 0 {
					parts = append(parts, &WordPart{BasePart{Subs: current}})
					current = nil
				}
			}
			current = append(current, s)
		}
		if len(current) > 0 {
			parts = append(parts, &WordPart{BasePart{Subs: current}})
		}
		return parts
	}
}
