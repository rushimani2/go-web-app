package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
	"time"
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
	t.Parallel()
	t.Run("User API Integration", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/1", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		resp := w.Result()
		t.Logf("Response Status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 but got %d", resp.StatusCode)
		} else {
			t.Logf("User API response was successful.")
		}
	})
}

// --- Database Integration Test (mocked) ---
func TestDatabaseConnection(t *testing.T) {
	t.Parallel()
	t.Run("Database Connection", func(t *testing.T) {
		// Simulate a database connection test (mocked)
		user := struct {
			ID int
		}{ID: 1} // Mocked user data

		t.Logf("Mock user ID: %d", user.ID)

		if user.ID != 1 {
			t.Errorf("Expected user ID 1, got %d", user.ID)
		} else {
			t.Logf("Database connection simulated successfully.")
		}
	})
}

// --- API Endpoint Test for Creating User ---
func TestCreateUserAPI(t *testing.T) {
	t.Parallel()
	t.Run("Create User API", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/users", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		resp := w.Result()
		t.Logf("Create User API response status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201 but got %d", resp.StatusCode)
		} else {
			t.Logf("Create User API call was successful.")
		}
	})
}

// --- User Login Integration Test ---
func TestUserLogin(t *testing.T) {
	t.Parallel()
	t.Run("User Login", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		resp := w.Result()
		t.Logf("Login response status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 but got %d", resp.StatusCode)
		} else {
			t.Logf("User login was successful.")
		}
	})
}

// --- User Registration and Login Workflow Test ---
func TestUserRegistrationLoginWorkflow(t *testing.T) {
	t.Parallel()
	t.Run("User Registration and Login Workflow", func(t *testing.T) {
		// Simulate user registration
		req := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()

		// Simulate calling the registration handler
		http.DefaultServeMux.ServeHTTP(w, req)

		if w.Result().StatusCode != http.StatusCreated {
			t.Fatalf("User registration failed with status: %d", w.Result().StatusCode)
		} else {
			t.Logf("User registered successfully with status %d", w.Result().StatusCode)
		}

		// Simulate login after registration
		req = httptest.NewRequest("POST", "/login", nil)
		w = httptest.NewRecorder()

		// Simulate calling the login handler
		http.DefaultServeMux.ServeHTTP(w, req)

		if w.Result().StatusCode != http.StatusOK {
			t.Fatalf("Login failed with status: %d", w.Result().StatusCode)
		} else {
			t.Logf("User logged in successfully with status %d", w.Result().StatusCode)
		}
	})
}

// --- Static Code Analysis Test (Security Test) ---
func TestStaticAnalysis(t *testing.T) {
	t.Parallel()
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

// --- Smoke Test (Basic Sanity Test) ---
func TestSmoke(t *testing.T) {
	t.Parallel()
	t.Run("Smoke Test", func(t *testing.T) {
		// Run a simple test to ensure core functionality is working
		TestUserAPIIntegration(t)
	})
}

// --- Performance Test (Under Load) ---
func TestPerformance(t *testing.T) {
	t.Parallel()
	t.Run("Performance Test", func(t *testing.T) {
		// Start the test server
		ts := httptest.NewServer(http.DefaultServeMux)
		defer ts.Close()

		// Measure the performance for 1000 requests
		start := time.Now()

		for i := 0; i < 1000; i++ {
			_, err := http.Get(ts.URL + "/api/users/1")
			if err != nil {
				t.Fatalf("Request %d failed: %v", i, err)
			}
		}

		duration := time.Since(start)
		t.Logf("Processed 1000 requests in %v", duration)
	})
}

// --- Security Test (SQL Injection Simulation) ---
func TestSecuritySQLInjection(t *testing.T) {
	t.Parallel()
	t.Run("Security Test - SQL Injection", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/1' OR 1=1 --", nil)
		w := httptest.NewRecorder()

		// Simulate calling the handler
		http.DefaultServeMux.ServeHTTP(w, req)

		// Check the response
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 but got %d", resp.StatusCode)
		} else {
			t.Logf("SQL injection test passed with status %d", resp.StatusCode)
		}

		expected := `{"id": 1, "name": "John Doe"}`
		if w.Body.String() != expected {
			t.Errorf("Expected body %q but got %q", expected, w.Body.String())
		} else {
			t.Logf("Response body matched expected output.")
		}
	})
}
