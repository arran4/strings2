package mappers

import (
	"bytes"
	"github.com/arran4/strings2/mappers/samples"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/arran4/strings2"
)

func TestMapAcronyms(t *testing.T) {
	words := []strings2.Word{
		nil,
		strings2.SingleCaseWord("user"),
		strings2.SingleCaseWord("id"),
	}

	mapper := MapAcronyms("ID")
	result := mapper(words)

	expected := []strings2.Word{
		nil,
		strings2.SingleCaseWord("user"),
		strings2.AcronymWord("ID"),
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMapAcronymsFromFile(t *testing.T) {
	f, err := os.CreateTemp("", "acronyms_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	_, _ = f.WriteString("HTML\nJSON\n")
	f.Close()

	mapper, err := MapAcronymsFromFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	words := []strings2.Word{
		strings2.SingleCaseWord("some"),
		strings2.SingleCaseWord("html"),
		strings2.SingleCaseWord("file"),
	}

	result := mapper(words)
	expected := []strings2.Word{
		strings2.SingleCaseWord("some"),
		strings2.AcronymWord("HTML"),
		strings2.SingleCaseWord("file"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Error case
	_, err = MapAcronymsFromFile("non_existent_file.txt")
	if err == nil {
		t.Error("Expected error for missing file")
	}
}

func TestMapAcronymsFromReader(t *testing.T) {
	reader := bytes.NewReader([]byte("HTML\nJSON\n"))
	mapper, err := MapAcronymsFromReader(reader)
	if err != nil {
		t.Fatal(err)
	}

	words := []strings2.Word{
		strings2.SingleCaseWord("some"),
		strings2.SingleCaseWord("html"),
	}

	result := mapper(words)
	expected := []strings2.Word{
		strings2.SingleCaseWord("some"),
		strings2.AcronymWord("HTML"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMapAcronymsFromBytesAndSamples(t *testing.T) {
	mapper, err := MapAcronymsFromBytes(samples.TechAcronyms)
	if err != nil {
		t.Fatal(err)
	}

	words := []strings2.Word{
		strings2.SingleCaseWord("my"),
		strings2.SingleCaseWord("api"),
	}

	result := mapper(words)
	expected := []strings2.Word{
		strings2.SingleCaseWord("my"),
		strings2.AcronymWord("API"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMapAcronymsFromURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("URL\nHTTP\n"))
	}))
	defer ts.Close()

	mapper, err := MapAcronymsFromURL(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	words := []strings2.Word{
		strings2.SingleCaseWord("my"),
		strings2.SingleCaseWord("url"),
	}

	result := mapper(words)
	expected := []strings2.Word{
		strings2.SingleCaseWord("my"),
		strings2.AcronymWord("URL"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
