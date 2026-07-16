package main

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func createChiRouter(logger *slog.Logger) chi.Router {
	return chi.NewRouter()
}
