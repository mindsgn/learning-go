package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeoutMiddlewareAddsDeadline(t *testing.T) {
	wrapped := Timeout(50)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deadline, ok := r.Context().Deadline()
		if !ok {
			t.Fatal("expected a deadline in the request context")
		}
		remaining := time.Until(deadline)
		if remaining <= 0 || remaining > 80*time.Millisecond {
			t.Fatalf("unexpected deadline remaining: %v", remaining)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	wrapped.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Fatalf("unexpected status: %d", rr.Code)
	}
}
