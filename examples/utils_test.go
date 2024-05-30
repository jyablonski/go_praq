// utils_test.go
package main

import (
	"testing"
)

func TestReturnStrChars(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected int
	}{
		{"test123", 7},
		{"", 0},
		{"hello, world", 12},
		{"1234567890", 10},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := return_str_chars(test.input)
			if result != test.expected {
				t.Errorf("Expected %d but got %d for input %q", test.expected, result, test.input)
			}
		})
	}
}
