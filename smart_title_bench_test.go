package strings2

import (
	"testing"
)

func runBench(b *testing.B, input string, fn SmartTitleAlgorithm) {
	words, _ := Parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(words, input)
	}
}

func BenchmarkSmartTitleAlgorithms(b *testing.B) {
	inputs := map[string]string{
		"Short":  "API",
		"Medium": "parse_HTTP_request",
		"Long":   "THE_QUICK_BROWN_FOX_JUMPS_OVER_THE_LAZY_DOG_API_CLIENT_HANDLER",
	}

	algos := map[string]SmartTitleAlgorithm{
		"Ratio":             algoRatio,
		"WholeSource":       algoWholeSource,
		"Structural":        algoStructural,
		"LexicalShape":      algoLexicalShape,
		"Scoring":           algoScoring,
		"StagedHybrid":      algoStagedHybrid,
	}

	for inputName, input := range inputs {
		b.Run(inputName, func(b *testing.B) {
			for algoName, fn := range algos {
				b.Run(algoName, func(b *testing.B) {
					runBench(b, input, fn)
				})
			}
		})
	}
}
