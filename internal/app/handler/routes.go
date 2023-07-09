package handler

import (
	"e-commerce/config"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(h Handler, cfg *config.Config) *chi.Mux {
	router := chi.NewMux()
	router.Use(middleware.Recoverer)
	router.Use(CORS)
	router.Use(middleware.Logger)
	router.Use(logger)
	return router
}
