package strings2

import (
	"testing"
)

func TestEdgeCases(t *testing.T) {
	// 1. Unicode in Mixed Case Splitting
	// Even though ExactCaseWord is a single word in the IL, OptionMixCaseSupport
	// instructs the formatter to split it based on casing.
	// This test verifies that this splitting works for both ASCII and Unicode.
	t.Run("Unicode MixCase", func(t *testing.T) {
		input := []Word{ExactCaseWord("helloWorld")} // ASCII mixed-case should be split
		res := ToFormattedCase(input, OptionMixCaseSupport(), OptionDelimiter("-"))
		if res != "hello-World" {
			t.Errorf("ASCII MixCase failed: got %q", res)
		}

		input = []Word{ExactCaseWord("helloWörld")} // Unicode Lower
		res = ToFormattedCase(input, OptionMixCaseSupport(), OptionDelimiter("-"))
		if res != "hello-Wörld" {
			// This is expected to work as W is ASCII
			t.Logf("Unicode MixCase 1: got %q", res)
		}

		input = []Word{ExactCaseWord("helloÖrld")} // Unicode Upper
		res = ToFormattedCase(input, OptionMixCaseSupport(), OptionDelimiter("-"))
		// Expectation: hello-Örld if unicode supported
		if res != "hello-Örld" {
			t.Errorf("Unicode MixCase 2: got %q, want %q", res, "hello-Örld")
		}
	})

	// 2. FirstUpperCaseWord with OptionMixCaseSupport
	t.Run("FirstUpperCaseWord MixCase", func(t *testing.T) {
		// FirstUpperCaseWord("camelCase") -> "Camelcase"
		// This test documents the behavior that FirstUpperCaseWord normalizes case,
		// preventing MixCase splitting from working as one might naively expect.
		input := []Word{FirstUpperCaseWord("camelCase")}
		res := ToFormattedCase(input, OptionMixCaseSupport(), OptionDelimiter("-"))

		// Current implementation behavior: "Camelcase"
		if res == "Camelcase" {
			// This is the current behavior.
		} else if res == "Camel-Case" {
			t.Log("FirstUpperCaseWord preserved internal case (Unexpected given current impl)")
		} else {
			t.Errorf("FirstUpperCaseWord behavior changed: got %q", res)
		}
	})

	// 3. Conflicting FirstUpper and FirstLower
	t.Run("Conflicting Options", func(t *testing.T) {
		input := []Word{SingleCaseWord("hello")}
		// FirstLower wins because it's applied last in ToFormattedCase
		res := ToFormattedCase(input, OptionFirstUpper(), OptionFirstLower())
		if res != "hello" {
			t.Errorf("Conflicting Options (FirstLower should win): got %q", res)
		}

		res = ToFormattedCase(input, OptionFirstLower(), OptionFirstUpper())
		if res != "hello" {
			t.Errorf("Conflicting Options (FirstLower should win regardless of order): got %q", res)
		}
	})

	// 4. Consecutive Uppercase in Mixed Case Splitting
	t.Run("Consecutive Uppercase", func(t *testing.T) {
		input := []Word{ExactCaseWord("JSONParser")}
		res := ToFormattedCase(input, OptionMixCaseSupport(), OptionDelimiter("-"))
		// Current simple implementation splits before every uppercase letter (except index 0)
		expected := "J-S-O-N-Parser"
		if res != expected {
			t.Errorf("Consecutive Uppercase: got %q, want %q", res, expected)
		}
	})

	// 5. UpperIndicator logic
	t.Run("UpperIndicator", func(t *testing.T) {
		input := []Word{SingleCaseWord("hello"), SingleCaseWord("world")}

		// Case A: Indicator equals Delimiter (Double Delimiter behavior)
		res := ToFormattedCase(input, OptionDelimiter("-"), OptionUpperIndicator("-"))
		if res != "hello--world" {
			t.Errorf("UpperIndicator Double Delimiter: got %q", res)
		}

		// Case B: Indicator != Delimiter
		// With proposed fix, this should use the indicator as delimiter
		res = ToFormattedCase(input, OptionDelimiter("-"), OptionUpperIndicator("="))
		if res != "hello=world" {
			t.Errorf("UpperIndicator Override: got %q, want %q", res, "hello=world")
		}
	})

	// 6. Empty and Nil Input
	t.Run("Empty and Nil Input", func(t *testing.T) {
		res := ToFormattedCase(nil)
		if res != "" {
			t.Errorf("Nil input: got %q, want empty string", res)
		}

		res = ToFormattedCase([]Word{})
		if res != "" {
			t.Errorf("Empty input: got %q, want empty string", res)
		}

		res = ToFormattedCase([]Word{SingleCaseWord("")})
		if res != "" {
			t.Errorf("Empty word: got %q, want empty string", res)
		}
	})
}
