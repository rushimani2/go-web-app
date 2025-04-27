package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	router := setupRouter() // assuming you have a setupRouter() or similar
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK but got %d", resp.StatusCode)
	}
}
