package strings2_test

import (
	"testing"

	"github.com/arran4/strings2"
)

func TestFormatPrecedence(t *testing.T) {
	words := []strings2.Word{strings2.ExactCaseWord("helloWorld")}

	// OptionMixCaseSupport is applied BEFORE screaming, screaming overwrites case.
	resScreaming := strings2.ToFormattedCase(words, strings2.OptionMixCaseSupport(), strings2.OptionDelimiter("-"), strings2.OptionCaseMode(strings2.CMScreaming))
	if resScreaming != "HELLO-WORLD" {
		t.Errorf("OptionMixCaseSupport with Screaming expected HELLO-WORLD, got %s", resScreaming)
	}

	resWhispering := strings2.ToFormattedCase(words, strings2.OptionMixCaseSupport(), strings2.OptionDelimiter("-"), strings2.OptionCaseMode(strings2.CMWhispering))
	if resWhispering != "hello-world" {
		t.Errorf("OptionMixCaseSupport with Whispering expected hello-world, got %s", resWhispering)
	}
}
