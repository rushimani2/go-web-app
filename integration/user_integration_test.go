package main

import (
	"testing"
)

// Unit Test Example
func TestExampleUnitFunction(t *testing.T) {
	result := 2 + 2
	expected := 4
	if result != expected {
		t.Errorf("Addition result was incorrect: got %d, want %d", result, expected)
	}
}
