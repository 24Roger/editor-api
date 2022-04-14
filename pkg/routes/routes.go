package routes

import (
	"github.com/24Roger/editor-api/pkg/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Route() *chi.Mux {
	mux := chi.NewMux()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler,
	)

	mux.Get("/all-projects", controllers.FindAllProjects)
	mux.Post("/new-project", controllers.NewProject)

	return mux
}
