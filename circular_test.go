package strings2_test

import (
	"testing"

	"github.com/arran4/strings2"
)

func TestCircularRestoration(t *testing.T) {
	// Test case: FromSentenceToSentence should restore original string
	input := "Hello to all the good-doers out there"

	subs, _ := strings2.StringToSubParts(input)
	delims := map[rune]bool{' ': true, '-': true}
	partitioner := strings2.NewPartitioner(strings2.PartitionerConfig{
		Delimiters:  delims,
		SplitCamel:  false, // Sentence usually space delimited
		PreserveSep: true,
	})

	parts := strings2.SubPartsToParts(subs, partitioner)
	words := strings2.PartsToWords(parts, nil) // This will produce SeparatorWords

	restored := strings2.ToFormattedCase(words, strings2.OptionDelimiter(""))

	if restored != input {
		t.Errorf("Circular restoration failed.\nInput:    %q\nRestored: %q", input, restored)
	}
}

func TestFromSentenceToSentence(t *testing.T) {
	input := "Hello to all the good-doers out there"

	delims := map[rune]bool{' ': true, '-': true}
	words, _ := strings2.Parse(input, strings2.PartitionerConfig{
		Delimiters:  delims,
		PreserveSep: true,
	})

	restored := strings2.ToFormattedCase(words, strings2.OptionDelimiter(""))

	if restored != input {
		t.Errorf("FromSentenceToSentence failed.\nInput:    %q\nRestored: %q", input, restored)
	}
}
