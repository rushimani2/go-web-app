package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIPPageIntegration(t *testing.T) {
	req := httptest.NewRequest("GET", "/ip", nil)
	w := httptest.NewRecorder()

	ipPage(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}

func TestHomePageIntegration(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	homePage(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}
}
