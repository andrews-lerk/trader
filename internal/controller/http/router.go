package http

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"trader/internal/controller/http/handlers/strategy"
	"trader/internal/controller/http/middleware"
)

func NewRouter(logger *slog.Logger) *chi.Mux {

	router := chi.NewRouter()

	// Setup middleware
	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.Recoverer)
	router.Use(middleware.Logger(logger))

	// Setup routes
	router.Post("/strategy/update", strategy.Update(logger))

	return router
}
