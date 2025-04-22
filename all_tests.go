// all_tests.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExampleUnitFunction(t *testing.T) { // Correct Test function
	result := 2 + 2
	expected := 4
	if result != expected {
		t.Errorf("Unit Test Failed: got %d, want %d", result, expected)
	}
}

func TestExampleHTTPHandlerIntegration(t *testing.T) { // Correct Test function
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Integration Test Response"))
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Integration Test Failed: status code %d", rr.Code)
	}
}

func BenchmarkExamplePerformance(b *testing.B) { // Correct Benchmark function
	for i := 0; i < b.N; i++ {
		_ = 2 + 2
	}
}
