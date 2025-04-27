// main_test.go
package main

import "testing"

func TestHomePage(t *testing.T) {
	// Call the HomePage handler directly, assuming it writes to the ResponseWriter
	// and does not require any complex HTTP request/response checks.

	err := HomePage(nil, nil)

	// If HomePage is supposed to return an error or do something, you can check it here.
	// For a simple "Hello World", you might expect it to not return any error.
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
