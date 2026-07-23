package strings2

import (
	"strings"
	"testing"
)

// SmartTitleAlgorithm interface for evaluating multiple distinct uppercase parsing logic.
type SmartTitleAlgorithm func(words []Word, input string) string

// Algo A: Ratio
func algoRatio(words []Word, input string) string {
	res, _ := ToTitleCase(words, OptionSmartTitleThreshold(func(wc int) float64 { return 0.5 }))
	return res
}

// Algo B: Whole Source
func algoWholeSource(words []Word, input string) string {
	hasLower := false
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			hasLower = true
			break
		}
	}

	if !hasLower && len(input) > 0 {
		res, _ := ToTitleCase(words, OptionSmartTitleUpperMode(SmartTitleUpperNormalize))
		return res
	}
	res, _ := ToTitleCase(words, OptionSmartTitleUpperMode(SmartTitleUpperPreserve))
	return res
}

// Algo C: Provenance
func algoProvenance(words []Word, input string) string {
	// Without modifying SmartAcronyms, they are all AcronymWords if length > 1
	// We'll mimic this by preserving AcronymWord and normalizing UpperCaseWord.
	res, _ := ToTitleCase(words, OptionSmartTitleUpperMode(SmartTitleUpperPreserve)) // Default ToTitleCase handles this poorly because of parsing
	return res
}

// Algo D: Structural
func algoStructural(words []Word, input string) string {
	// The default auto mode includes structural checks.
	res, _ := ToTitleCase(words, OptionSmartTitleUpperMode(SmartTitleUpperAuto))
	return res
}

// Algo E: Lexical Shape
func algoLexicalShape(words []Word, input string) string {
	res, _ := ToTitleCase(words, OptionSmartTitleAcronymPredicate(func(w string) bool {
		// Vowel presence check (inverted: if no vowels, likely acronym, though crude)
		vowels := "AEIOUY"
		hasVowel := strings.ContainsAny(w, vowels)
		return !hasVowel || len(w) == 1
	}))
	return res
}

// Algo F: Acronym Predicate
func algoAcronymPredicate(words []Word, input string) string {
	res, _ := ToTitleCase(words, OptionSmartTitleAcronymPredicate(func(w string) bool {
		switch w {
		case "NASA", "API", "HTTP", "URL", "COVID":
			return true
		}
		return false
	}))
	return res
}

// Algo G: Explicit Mode
func algoExplicitMode(words []Word, input string) string {
	res, _ := ToTitleCase(words, OptionSmartTitleUpperMode(SmartTitleUpperNormalize))
	return res
}

// Algo H: Multi-Signal Scoring
func algoScoring(words []Word, input string) string {
	// Very simple scoring logic outside of the framework
	res, _ := ToTitleCase(words, OptionSmartTitleAcronymPredicate(func(w string) bool {
		score := 0
		if len(w) <= 4 { score++ }
		if w == "HTTP" || w == "API" { score += 2 }
		return score >= 2
	}))
	return res
}

// Algo I: Staged Hybrid (Preferred)
func algoStagedHybrid(words []Word, input string) string {
	// Uses the new default behavior with the auto upper mode implicitly.
	res, _ := ToTitleCase(words)
	return res
}

func TestSmartTitleAlgorithms(t *testing.T) {
	tests := []string{
		"the lord of the rings",
		"A_NEW_HOPE",
		"THE_LORD_OF_THE_RINGS",
		"parse_HTTP_request",
		"HTTP_request",
		"mixed-UP-Kebab",
		"NASA_API_CLIENT",
		"COVID_19_RESPONSE",
		"A_B_TESTING",
		"API",
		"GO_TO_URL",
		"ONE_TWO",
		"ONE_TWO_THREE_FOUR",
		"One_TWO_THREE_Four",
		"XMLHttpRequest",
		"JSON_API_response",
		"user_ID",
		"ID_user",
		"version_2_API",
		"single",
		"I",
		"",
		"_",
		"123",
		"café_BISTRO", // Unicode test
	}

	algos := map[string]SmartTitleAlgorithm{
		"Ratio":             algoRatio,
		"WholeSource":       algoWholeSource,
		"Provenance":        algoProvenance,
		"Structural":        algoStructural,
		"LexicalShape":      algoLexicalShape,
		"AcronymPredicate":  algoAcronymPredicate,
		"ExplicitMode":      algoExplicitMode,
		"Scoring":           algoScoring,
		"StagedHybrid":      algoStagedHybrid,
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			words, _ := Parse(input)
			t.Logf("Input: %q", input)
			for name, fn := range algos {
				out := fn(words, input)
				t.Logf("  %16s : %q", name, out)
			}
		})
	}
}
