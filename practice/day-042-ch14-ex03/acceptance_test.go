package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	elog "learninggo/practice/day-042-ch14-ex03/log"
	"learninggo/practice/internal/testkit"
)

func TestContextHelpers(t *testing.T) {
	ctx := elog.ContextWithLevel(context.Background(), elog.Debug)
	level, ok := elog.LevelFromContext(ctx)
	if !ok || level != elog.Debug {
		t.Fatalf("unexpected level: %v, %v", level, ok)
	}
}

func TestLogAndMiddleware(t *testing.T) {
	output := testkit.CaptureStdout(t, func() {
		ctx := elog.ContextWithLevel(context.Background(), elog.Debug)
		elog.Log(ctx, elog.Debug, "debug message")
	})
	if testkit.NormalizeOutput(output) != "debug message" {
		t.Fatalf("unexpected log output: %q", output)
	}
	var seen bool
	h := elog.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level, ok := elog.LevelFromContext(r.Context())
		if !ok || level != elog.Info {
			t.Fatalf("middleware should inject info level, got %v %v", level, ok)
		}
		seen = true
		w.WriteHeader(http.StatusNoContent)
	}))
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/?log_level=info", nil)
	h.ServeHTTP(rr, req)
	if !seen {
		t.Fatal("expected wrapped handler to run")
	}
}
