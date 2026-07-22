package mappers

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/arran4/strings2"
)

// MapLowercase returns a WordMapper that converts matching words (case-insensitive) into LowercaseWord (lowercase).
// It acts as a predicate for minor words (prepositions, conjunctions) to stay lowercase during title casing.
func MapLowercase(wordsToKeepLower ...string) func([]strings2.Word) []strings2.Word {
	if len(wordsToKeepLower) == 0 {
		return func(words []strings2.Word) []strings2.Word { return words }
	}

	lowerMap := make(map[string]struct{}, len(wordsToKeepLower))
	for _, a := range wordsToKeepLower {
		lowerMap[strings.ToLower(a)] = struct{}{}
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
			if _, ok := lowerMap[lowerS]; ok {
				result = append(result, strings2.LowercaseWord(lowerS))
			} else {
				result = append(result, w)
			}
		}
		return result
	}
}

// MapLowercaseFromReader reads lowercase predicates from an io.Reader (one per line) and returns a WordMapper.
func MapLowercaseFromReader(r io.Reader) (func([]strings2.Word) []strings2.Word, error) {
	var predicates []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			predicates = append(predicates, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return MapLowercase(predicates...), nil
}

// MapLowercaseFromBytes reads lowercase predicates from a []byte (one per line) and returns a WordMapper.
func MapLowercaseFromBytes(b []byte) (func([]strings2.Word) []strings2.Word, error) {
	return MapLowercaseFromReader(bytes.NewReader(b))
}

// MapLowercaseFromFile reads a file containing one lowercase predicate per line and returns a WordMapper.
func MapLowercaseFromFile(filename string) (func([]strings2.Word) []strings2.Word, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return MapLowercaseFromReader(file)
}

// MapLowercaseFromURL fetches lowercase predicates from a given URL (one per line) and returns a WordMapper.
func MapLowercaseFromURL(url string) (func([]strings2.Word) []strings2.Word, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to fetch lowercase words from URL: %s, status code: %d", url, resp.StatusCode)
	}
	return MapLowercaseFromReader(resp.Body)
}
