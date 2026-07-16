package main

import (
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"
)

func createChiRouter(logger *slog.Logger) chi.Router {
	return chi.NewRouter()
}

func buildText(now time.Time) string {
	return ""
}

func buildJSON(now time.Time) string {
	return "{}"
}
