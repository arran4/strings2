package mappers

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
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
		result := make([]strings2.Word, 0, len(words))
		for _, w := range words {
			if w == nil {
				result = append(result, nil)
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

// MapAcronymsFromReader reads acronyms from an io.Reader (one per line) and returns a WordMapper.
func MapAcronymsFromReader(r io.Reader) (func([]strings2.Word) []strings2.Word, error) {
	var acronyms []string
	scanner := bufio.NewScanner(r)
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

// MapAcronymsFromBytes reads acronyms from a []byte (one per line) and returns a WordMapper.
func MapAcronymsFromBytes(b []byte) (func([]strings2.Word) []strings2.Word, error) {
	return MapAcronymsFromReader(bytes.NewReader(b))
}

// MapAcronymsFromFile reads a file containing one acronym per line and returns a WordMapper.
func MapAcronymsFromFile(filename string) (func([]strings2.Word) []strings2.Word, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return MapAcronymsFromReader(file)
}

// MapAcronymsFromURL fetches acronyms from a given URL (one per line) and returns a WordMapper.
func MapAcronymsFromURL(url string) (func([]strings2.Word) []strings2.Word, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return MapAcronymsFromReader(resp.Body)
}
