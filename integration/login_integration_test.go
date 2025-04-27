// login_integration_test.go
package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
