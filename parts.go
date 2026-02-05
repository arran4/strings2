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
	var parts []Part
	var current []SubPart

	for _, s := range subs {
		if s.Rune() == delim {
			if len(current) > 0 {
				parts = append(parts, &WordPart{BasePart{Subs: current}})
				current = nil
			}
			// We usually discard the delimiter in casing conversions, or make it a SeparatorPart?
			// For "ToWords", we usually want the content.
			// Let's discard for now, or maybe make it configurable?
			// The user said "SubParts... count of things such as SpaceDelimiterCount...".
			// But for "Parts", usually we want the words.
		} else {
			current = append(current, s)
		}
	}
	if len(current) > 0 {
		parts = append(parts, &WordPart{BasePart{Subs: current}})
	}
	return parts
}

// CamelCasePartitioner splits on case transitions.
func CamelCasePartitioner(subs []SubPart) []Part {
	var parts []Part
	var current []SubPart

	for i, s := range subs {
		// Transition check
		isSplit := false
		if i > 0 {
			prev := subs[i-1]

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
