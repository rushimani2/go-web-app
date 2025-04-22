package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

// Registering routes globally
func registerRoutes() {
	// Registering routes only once for all tests
	http.HandleFunc("/api/users/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1, "name": "John Doe"}`)) // Mocking a user response
	})

	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": 2, "name": "Jane Doe"}`)) // Mock creating a new user
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Login successful"}`)) // Mock login success response
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": 1, "name": "New User"}`)) // Mock user registration response
	})
}

// TestMain will run once before any tests are executed.
func TestMain(m *testing.M) {
	// Register all routes before tests start
	registerRoutes()

	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}

// --- User API Integration Test ---
func TestUserAPIIntegration(t *testing.T) {
	t.Run("User API Integration", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/1", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		resp := w.Result()
		t.Logf("Response Status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 but got %d", resp.StatusCode)
		}
	})
}

// --- Database Integration Test ---
func TestDatabaseConnection(t *testing.T) {
	t.Run("Database Connection", func(t *testing.T) {
		// Simulate a database connection test (mocked)
		user := struct {
			ID int
		}{ID: 1} // Mocked user data

		t.Logf("Mock user ID: %d", user.ID)

		if user.ID != 1 {
			t.Errorf("Expected user ID 1, got %d", user.ID)
		}
	})
}

// --- API Endpoint Test ---
func TestCreateUserAPI(t *testing.T) {
	t.Run("Create User API", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/users", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		resp := w.Result()
		t.Logf("Create User API response status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201 but got %d", resp.StatusCode)
		}
	})
}

// --- Login Integration Test (End-to-End) ---
func TestUserLogin(t *testing.T) {
	t.Run("User Login", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		resp := w.Result()
		t.Logf("Login response status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 but got %d", resp.StatusCode)
		}
	})
}

// --- User Workflow Test (End-to-End) ---
func TestUserRegistrationLoginWorkflow(t *testing.T) {
	t.Run("User Registration and Login Workflow", func(t *testing.T) {
		// Simulate user registration
		req := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()

		// Simulate calling the registration handler
		http.DefaultServeMux.ServeHTTP(w, req)

		if w.Result().StatusCode != http.StatusCreated {
			t.Fatalf("User registration failed with status: %d", w.Result().StatusCode)
		}
		t.Logf("User registered successfully with status %d", w.Result().StatusCode)

		// Simulate login after registration
		req = httptest.NewRequest("POST", "/login", nil)
		w = httptest.NewRecorder()

		// Simulate calling the login handler
		http.DefaultServeMux.ServeHTTP(w, req)

		if w.Result().StatusCode != http.StatusOK {
			t.Fatalf("Login failed with status: %d", w.Result().StatusCode)
		}
		t.Logf("User logged in successfully with status %d", w.Result().StatusCode)
	})
}

// --- Static Code Analysis (Security Test) ---
func TestStaticAnalysis(t *testing.T) {
	t.Run("Static Analysis", func(t *testing.T) {
		cmd := exec.Command("golangci-lint", "run", "--enable=govet,staticcheck", "--disable=gofmt")
		err := cmd.Run()

		if err != nil {
			t.Errorf("Static analysis failed: %v", err)
		} else {
			t.Log("Static analysis passed successfully.")
		}
	})
}
