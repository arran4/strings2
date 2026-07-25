package strings2

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

// OptionAcronymList returns a WordMapper that converts matching words (case-insensitive) into AcronymWord.
// The provided acronyms list dictates the exact case the AcronymWord will take.
func OptionAcronymList(acronyms []string) WordMapper {
	if len(acronyms) == 0 {
		return func(words []Word) []Word { return words }
	}

	acronymMap := make(map[string]string, len(acronyms))
	for _, a := range acronyms {
		acronymMap[strings.ToLower(a)] = a
		// Map the exact case back to itself to be consistent with lowercase mappings
		acronymMap[a] = a
	}

	return func(words []Word) []Word {
		result := make([]Word, 0, len(words))
		for _, w := range words {
			if w == nil {
				result = append(result, nil)
				continue
			}
			s := w.String()

			// Attempt to match lowercase representation
			lowerS := strings.ToLower(s)
			if replacement, ok := acronymMap[lowerS]; ok {
				result = append(result, AcronymWord(replacement))
			} else {
				// We also check exact original case match just in case
				if replacement, ok := acronymMap[s]; ok {
					result = append(result, AcronymWord(replacement))
				} else {
					result = append(result, w)
				}
			}
		}
		return result
	}
}

// LoadAcronymsFromReader reads acronyms from an io.Reader (one per line) and returns a WordMapper.
func LoadAcronymsFromReader(r io.Reader) (WordMapper, error) {
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

	return OptionAcronymList(acronyms), nil
}

// LoadAcronymsFromFile reads a file containing one acronym per line and returns a WordMapper.
func LoadAcronymsFromFile(filepath string) (WordMapper, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return LoadAcronymsFromReader(file)
}

// LoadAcronymsFromURL fetches acronyms from a given URL (one per line) and returns a WordMapper.
func LoadAcronymsFromURL(url string) (WordMapper, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return LoadAcronymsFromReader(resp.Body)
}
