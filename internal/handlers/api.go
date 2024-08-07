package handlers

import (
	"github.com/Lazer430/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chiMiddle "github.com/go-chi/chi/middleware"
)

// Define the handler function
func Handler(r *chi.Mux) {
	//Add global middleware
	r.Use(chiMiddle.StripSlashes)

	// Define Route
	r.Route("/account", func(router chi.Router) {

		// Auth middleware
		router.Use(middleware.Authorization)

		// Get method
		router.Get("/coins", GetCoinBalance)
	})

}

func RegisterHandlers(r *chi.Mux) {
	Handler(r)
}
