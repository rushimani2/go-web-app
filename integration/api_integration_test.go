package integration

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"your_project/internal/service"
	"your_project/internal/database"
)

// --- User API Integration Test ---
func TestUserAPIIntegration(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/users/1", nil)
	w := httptest.NewRecorder()

	service.HandleUserRequest(w, req) // Replace with your actual handler

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}

// --- Database Integration Test ---
func TestDatabaseConnection(t *testing.T) {
	db, err := database.Connect("localhost:5432")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	user, err := db.GetUserByID(1)
	if err != nil {
		t.Fatalf("Error fetching user: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", user.ID)
	}
}

// --- API Endpoint Test ---
func TestCreateUserAPI(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/users", nil) // Add JSON body if needed
	w := httptest.NewRecorder()

	service.CreateUserHandler(w, req) // Replace with actual handler

	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201 but got %d", resp.StatusCode)
	}
}

// --- Login Integration Test (End-to-End) ---
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

// --- User Workflow Test (End-to-End) ---
func TestUserRegistrationLoginWorkflow(t *testing.T) {
	app := app.New()

	// Simulate user registration
	req := httptest.NewRequest("POST", "/register", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusCreated {
		t.Fatalf("User registration failed with status: %d", w.Result().StatusCode)
	}

	// Simulate login
	req = httptest.NewRequest("POST", "/login", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("Login failed with status: %d", w.Result().StatusCode)
	}
}

// --- Performance Load Test ---
func BenchmarkUserAPI(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		httptest.NewRecorder()
		handleRequest(w, req) // Your handler function
	}
}

// --- Health Check Smoke Test ---
func TestHealthCheck(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("Health check failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}

// --- Canary Test ---
func TestCanaryDeployment(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/canary")
	if err != nil {
		t.Fatalf("Canary check failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}

// --- Static Code Analysis (Security Test) ---
func TestStaticAnalysis(t *testing.T) {
	cmd := exec.Command("golangci-lint", "run", "--enable=govet,staticcheck", "--disable=gofmt")
	err := cmd.Run()
	if err != nil {
		t.Errorf("Static analysis failed: %v", err)
	}
}
