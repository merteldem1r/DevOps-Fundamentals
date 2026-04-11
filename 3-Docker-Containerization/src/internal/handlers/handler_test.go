package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGlobalHandler_GetHealth(t *testing.T) {
	h := NewGlobalHandler("<3")

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

	if res.Msg != "OK" {
		t.Fatalf("got msg %q, want %q", res.Msg, "OK")
	}
}

func TestGlobalHandler_Get(t *testing.T) {
	h := NewGlobalHandler("<3")

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

	if res.Msg != "<3" {
		t.Fatalf("got msg %q, want %q", res.Msg, "<3")
	}
}
