package routes

import (
	"github.com/BerkAkipek/go-two-endpoints/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/api", func(api chi.Router) {
		api.Get("/users", handlers.GetUsersHandler)
		api.Post("/users", handlers.CreateUserHandler)
	})

	return r
}
