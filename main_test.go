// main_test.go

package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"your_project/internal/app" // Replace with the actual path to your app package
)

// Define the setupRouter function
func setupRouter() *http.ServeMux {
	// Replace with actual router setup
	mux := http.NewServeMux()
	// Add your route handlers here
	mux.HandleFunc("/login", app.LoginHandler) // Example
	return mux
}

func TestLoginPage(t *testing.T) {
	req := httptest.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()
	router := setupRouter()

	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}
