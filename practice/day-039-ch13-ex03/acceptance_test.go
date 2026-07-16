package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBuildTextAndJSON(t *testing.T) {
	now := time.Date(2026, time.July, 16, 12, 34, 56, 0, time.UTC)
	if got := buildText(now); got != now.Format(time.RFC3339) {
		t.Fatalf("buildText mismatch: %q", got)
	}
	payload := map[string]any{}
	if err := json.Unmarshal([]byte(buildJSON(now)), &payload); err != nil {
		t.Fatalf("buildJSON returned invalid JSON: %v", err)
	}
	if payload["day_of_week"] != "Thursday" {
		t.Fatalf("unexpected day_of_week: %v", payload["day_of_week"])
	}
}

func TestRouterNegotiatesJSON(t *testing.T) {
	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	router := createChiRouter(logger)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept", "application/json")
	req.RemoteAddr = "127.0.0.1:4567"
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rr.Code)
	}
	payload := map[string]any{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected JSON response, got %q: %v", rr.Body.String(), err)
	}
}
