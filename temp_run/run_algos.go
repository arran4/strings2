package main

import (
	"fmt"
	"github.com/arran4/strings2"
)

func main() {
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
		"café_BISTRO",
	}

	algos := []struct {
		name string
		fn   func(words []strings2.Word, input string) string
	}{
		{"Ratio", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleThreshold(func(wc int) float64 { return 0.5 }))
			return res
		}},
		{"WholeSource", func(words []strings2.Word, input string) string {
			hasLower := false
			for _, r := range input {
				if r >= 'a' && r <= 'z' {
					hasLower = true
					break
				}
			}
			if !hasLower && len(input) > 0 {
				res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleUpperMode(strings2.SmartTitleUpperNormalize))
				return res
			}
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleUpperMode(strings2.SmartTitleUpperPreserve))
			return res
		}},
		{"Provenance", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleUpperMode(strings2.SmartTitleUpperPreserve))
			return res
		}},
		{"Structural", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleUpperMode(strings2.SmartTitleUpperAuto))
			return res
		}},
		{"LexicalShape", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleAcronymPredicate(func(w string) bool {
				vowels := "AEIOUY"
				hasVowel := false
				for _, v := range vowels {
					for _, c := range w {
						if c == v {
							hasVowel = true
							break
						}
					}
				}
				return !hasVowel || len(w) == 1
			}))
			return res
		}},
		{"AcronymPredicate", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleAcronymPredicate(func(w string) bool {
				switch w {
				case "NASA", "API", "HTTP", "URL", "COVID":
					return true
				}
				return false
			}))
			return res
		}},
		{"ExplicitMode", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleUpperMode(strings2.SmartTitleUpperNormalize))
			return res
		}},
		{"Scoring", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words, strings2.OptionSmartTitleAcronymPredicate(func(w string) bool {
				score := 0
				if len(w) <= 4 { score++ }
				if w == "HTTP" || w == "API" { score += 2 }
				return score >= 2
			}))
			return res
		}},
		{"StagedHybrid", func(words []strings2.Word, input string) string {
			res, _ := strings2.ToTitleCase(words)
			return res
		}},
	}

	for _, t := range tests {
		fmt.Printf("\t\t{input: %q, expected: map[string]string{\n", t)
		words, _ := strings2.Parse(t)
		for _, a := range algos {
			fmt.Printf("\t\t\t%q: %q,\n", a.name, a.fn(words, t))
		}
		fmt.Printf("\t\t}},\n")
	}
}
