// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest("GET", "/home", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HomePage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", rr.Code)
	}
}
