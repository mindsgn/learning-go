package main

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateChiRouterWithLogging(t *testing.T) {
	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	router := createChiRouter(logger)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:4567"
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rr.Code)
	}
	if _, err := time.Parse(time.RFC3339, rr.Body.String()); err != nil {
		t.Fatalf("expected RFC3339 time, got %q: %v", rr.Body.String(), err)
	}
	if !bytes.Contains(logs.Bytes(), []byte("127.0.0.1")) {
		t.Fatalf("expected log output to include the remote IP, got %q", logs.String())
	}
}
