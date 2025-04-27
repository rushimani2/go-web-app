package main

import "testing"

// Sample function to test
func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    
    if result != 3 {
        t.Errorf("Expected 3, but got %d", result)
    }
}
