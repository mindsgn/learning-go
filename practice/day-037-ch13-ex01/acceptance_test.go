package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateServeMux(t *testing.T) {
	h := createServeMux()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rr.Code)
	}
	if _, err := time.Parse(time.RFC3339, rr.Body.String()); err != nil {
		t.Fatalf("expected RFC3339 time, got %q: %v", rr.Body.String(), err)
	}
	badReq := httptest.NewRequest(http.MethodPost, "/", nil)
	badRR := httptest.NewRecorder()
	h.ServeHTTP(badRR, badReq)
	if badRR.Code != http.StatusMethodNotAllowed {
		t.Fatalf("POST should return 405, got %d", badRR.Code)
	}
}
