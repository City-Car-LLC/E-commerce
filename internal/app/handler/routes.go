package handler

import (
	"e-commerce/config"

	_ "e-commerce/cmd/e-commerce/docs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter godoc
// @title E-COMMERCE API
// @version 0.1
// @contact.name API Support
// @contact.url https://t.me/farafetch
// @contact.email faranush.karimov@gmail.com
// @license.name Custom License
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func NewRouter(h Handler, cfg *config.Config) *chi.Mux {
	router := chi.NewMux()
	router.Use(middleware.Recoverer)
	router.Use(CORS)
	router.Use(middleware.Logger)
	router.Use(logger)

	router.Get("/swagger/*", httpSwagger.Handler())

	api := chi.NewRouter()
	router.Mount("/api", api)
	api.Route("/categories", func(r chi.Router) {
		r.Post("/", H(h.CreateCategories))
		r.Get("/", H(h.ReadCategories))
	})

	return router
}
