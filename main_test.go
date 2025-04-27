// main_test.go
package main

import "testing"

func TestGenerateGreeting(t *testing.T) {
    name := "World"
    expected := "Hello, World!"

    result := GenerateGreeting(name)
    
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}
