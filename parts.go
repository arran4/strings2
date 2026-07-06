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


// DelimiterDetector detects if a delimiter starts at the given index in subs.
// It returns the length of the delimiter in subparts (runes), or 0 if no delimiter is found.
type DelimiterDetector func(subs []SubPart, index int) (length int)

type PartitionerConfig struct {
	Delimiters        map[rune]bool
	DelimiterDetector DelimiterDetector
	SplitCamel        bool

	NumberMode  NumberMode
	PreserveSep bool // If true, delimiters are returned as SeparatorPart instead of discarded
}

// NewPartitioner creates a partitioner with specific configuration.
func NewPartitioner(cfg PartitionerConfig) Partitioner {
	return func(subs []SubPart) []Part {
		var parts []Part
		var current []SubPart

		for i := 0; i < len(subs); {
			s := subs[i]

			delimLen := 0
			if cfg.DelimiterDetector != nil {
				delimLen = cfg.DelimiterDetector(subs, i)
				if delimLen < 0 {
					delimLen = 0
				}
				if delimLen > len(subs)-i {
					delimLen = len(subs) - i
				}
			}
			if delimLen == 0 && cfg.Delimiters != nil && cfg.Delimiters[s.Rune()] {
				delimLen = 1
			}

			// Check if current rune(s) is a delimiter
			if delimLen > 0 {
				if len(current) > 0 {
					parts = append(parts, &WordPart{BasePart{Subs: current}})
					current = nil
				}
				if cfg.PreserveSep {
					parts = append(parts, &SeparatorPart{BasePart{Subs: subs[i : i+delimLen]}})
				}
				i += delimLen
				continue
			}

			// Transition check
			isSplit := false
			if (cfg.SplitCamel || cfg.NumberMode != NumberModeNone) && len(current) > 0 {
				prev := current[len(current)-1]
				// Note: if prev was delimiter, current is empty or started anew.
				// We rely on current being non-empty to check transitions within a word chunk.

				if cfg.SplitCamel {
					isPrevLower := prev.IsLower()
					isPrevUpper := prev.IsUpper()
					isCurrUpper := s.IsUpper()

					if cfg.NumberMode == NumberModeTreatAsLowercase {
						if prev.IsDigit() {
							isPrevLower = true
						}
					}

					// lower -> Upper
					if isPrevLower && isCurrUpper {
						isSplit = true
					}

					// Upper -> Upper -> lower (PDFLoader split at L)
					if i+1 < len(subs) {
						next := subs[i+1]
						isNextLower := next.IsLower()
						if cfg.NumberMode == NumberModeTreatAsLowercase && next.IsDigit() {
							isNextLower = true
						}
						if isPrevUpper && isCurrUpper && isNextLower {
							isSplit = true
						}
					}

					// MergeRecursive specific rule: digit -> Upper triggers a split, similar to lower -> Upper
					if cfg.NumberMode == NumberModeMergeWithWord {
						if prev.IsDigit() && isCurrUpper {
							isSplit = true
						}
					}
				}

				if cfg.NumberMode == NumberModeSplitAlways {
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
			i++
		}
		if len(current) > 0 {
			parts = append(parts, &WordPart{BasePart{Subs: current}})
		}
		return parts
	}
}
