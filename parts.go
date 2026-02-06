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
	return NewPartitioner(PartitionerConfig{
		Delimiters: map[rune]bool{delim: true},
	})(subs)
}

// CamelCasePartitioner splits on case transitions.
func CamelCasePartitioner(subs []SubPart) []Part {
	return NewPartitioner(PartitionerConfig{
		SplitCamel: true,
	})(subs)
}

type PartitionerConfig struct {
	Delimiters  map[rune]bool
	SplitCamel  bool
	SplitNumber bool
	PreserveSep bool // If true, delimiters are returned as SeparatorPart instead of discarded
}

// NewPartitioner creates a partitioner with specific configuration.
func NewPartitioner(cfg PartitionerConfig) Partitioner {
	return func(subs []SubPart) []Part {
		var parts []Part
		var current []SubPart

		for i, s := range subs {
			// Check if current rune is a delimiter
			if cfg.Delimiters != nil && cfg.Delimiters[s.Rune()] {
				if len(current) > 0 {
					parts = append(parts, &WordPart{BasePart{Subs: current}})
					current = nil
				}
				if cfg.PreserveSep {
					parts = append(parts, &SeparatorPart{BasePart{Subs: []SubPart{s}}})
				}
				continue
			}

			// Transition check
			isSplit := false
			if (cfg.SplitCamel || cfg.SplitNumber) && i > 0 && len(current) > 0 {
				prev := subs[i-1]
				// Note: if prev was delimiter, current is empty or started anew.
				// We rely on current being non-empty to check transitions within a word chunk.

				if cfg.SplitCamel {
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

				if cfg.SplitNumber {
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
