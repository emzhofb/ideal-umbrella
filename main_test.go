// main_test.go
package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty name", "", "Hello, World!"},
		{"Non-empty name", "Indonesia", "Hello, Indonesia!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
