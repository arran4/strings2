package strings2

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		// TODO add comprehensive tests
	}

	// Test function
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCamelCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		// TODO add comprehensive tests
	}

	// Test function
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPascalCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		// TODO add comprehensive tests
	}

	// Test function
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToKebabCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		words    []Word
		options  []Option
		expected string
	}{
		// TODO add comprehensive tests
	}

	// Test function
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnakeCase(tt.words, tt.options...)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
