package mappers

import (
	"bufio"
	"os"
	"strings"

	"github.com/arran4/strings2"
)

// MapAcronyms returns a WordMapper that converts matching words (case-insensitive) into AcronymWord.
// The provided acronyms list dictates the exact case the AcronymWord will take.
// For example, if "ID" is provided, any word matching "id" (like "Id", "id", "ID") will become AcronymWord("ID").
func MapAcronyms(acronyms ...string) func([]strings2.Word) []strings2.Word {
	if len(acronyms) == 0 {
		return func(words []strings2.Word) []strings2.Word { return words }
	}

	acronymMap := make(map[string]string, len(acronyms))
	for _, a := range acronyms {
		acronymMap[strings.ToLower(a)] = a
	}

	return func(words []strings2.Word) []strings2.Word {
		var result []strings2.Word
		for _, w := range words {
			if w == nil {
				continue
			}
			s := w.String()
			lowerS := strings.ToLower(s)
			if replacement, ok := acronymMap[lowerS]; ok {
				result = append(result, strings2.AcronymWord(replacement))
			} else {
				result = append(result, w)
			}
		}
		return result
	}
}

// MapAcronymsFromFile reads a file containing one acronym per line and returns a WordMapper.
func MapAcronymsFromFile(filename string) (func([]strings2.Word) []strings2.Word, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var acronyms []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			acronyms = append(acronyms, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return MapAcronyms(acronyms...), nil
}
