package integration

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"your_project/internal/service"
)

func TestUserAPIIntegration(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/users/1", nil)
	w := httptest.NewRecorder()

	service.HandleUserRequest(w, req) // Replace with your actual handler

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}
