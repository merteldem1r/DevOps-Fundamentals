package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
)

func TestGlobalHandler_GetHealth(t *testing.T) {
	h := NewGlobalHandler(&config.Config{Message: "<3"}, nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rr := httptest.NewRecorder()

	h.GetHealth(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("got status %d, want %d", rr.Code, http.StatusOK)
	}

	if got := rr.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("got Content-Type %q, want %q", got, "application/json")
	}

	var res HandlerResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if res.Status != "Success" {
		t.Fatalf("got status %q, want %q", res.Status, "Success")
	}

	if res.Data != "OK" {
		t.Fatalf("got Data %q, want %q", res.Data, "OK")
	}
}

func TestGlobalHandler_Get(t *testing.T) {
	h := NewGlobalHandler(&config.Config{Message: "<3"}, nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rr := httptest.NewRecorder()

	h.Get(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("got status %d, want %d", rr.Code, http.StatusOK)
	}

	var res HandlerResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if res.Status != "Success" {
		t.Fatalf("got status %q, want %q", res.Status, "Success")
	}

	if res.Data != "<3" {
		t.Fatalf("got Data %q, want %q", res.Data, "<3")
	}
}
