package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginPageIntegration(t *testing.T) {
	req := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	loginPage(w, req) // Directly calling your loginPage handler

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}
