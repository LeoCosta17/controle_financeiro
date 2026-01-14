package main

import (
	"app/internal/handlers"
	"app/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LoadRoutes(s services.Services) *chi.Mux {

	r := chi.NewRouter()

	handlers := handlers.NewHandlers(s)

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/health", handlers.Health.APIHealth)
	r.Post("/users", handlers.Users.Create)
	r.Get("/users", handlers.Users.GetAll)

	return r
}
