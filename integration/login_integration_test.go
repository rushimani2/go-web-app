package integration

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"your_project/internal/app"
)

func TestUserLogin(t *testing.T) {
	app := app.New()

	req := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}
