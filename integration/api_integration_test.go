package integration

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"your_project/internal/service"
)

func TestCreateUserAPI(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/users", nil) // Add JSON body if needed
	w := httptest.NewRecorder()

	service.CreateUserHandler(w, req) // Replace with actual handler

	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201 but got %d", resp.StatusCode)
	}
}
