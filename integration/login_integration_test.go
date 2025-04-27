// login_integration_test.go

package integration

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"your_project/internal/app" // Replace with the actual path to your app package
)

// Define the loginPage variable (or function if needed)
var loginPage = "/login" // This could be your endpoint or the handler's path.

func TestLoginPageIntegration(t *testing.T) {
	req := httptest.NewRequest("POST", loginPage, nil) // Use the defined loginPage
	w := httptest.NewRecorder()
	router := app.SetupRouter() // Example, replace with your actual router setup

	router.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}
