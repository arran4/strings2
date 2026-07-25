package strings2_test

import (
	"os"
	"testing"
	"path/filepath"

	"github.com/arran4/strings2"
)

func TestOptionAcronymList(t *testing.T) {
	input := "Http Json Config"
	expected := "HTTP_JSON_Config" // Since AcronymWords are preserved in original case, they will not be lowercased by snake case

	result, err := strings2.ToSnake(input, strings2.OptionAcronymList([]string{"HTTP", "JSON"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestLoadAcronymsFromFile(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "acronyms.txt")
	err := os.WriteFile(file, []byte("HTTP\nJSON\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	mapper, err := strings2.LoadAcronymsFromFile(file)
	if err != nil {
		t.Fatalf("unexpected error loading acronyms: %v", err)
	}

	input := "Http Json Config"
	expected := "HTTP_JSON_Config"

	result, err := strings2.ToSnake(input, mapper)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
