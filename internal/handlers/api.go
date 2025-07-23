package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Mahjabin08/goapi/internal/middleware"
)
// starting a func name with upper case tells the compiler that it can be imported in other packages
//middleware is basically called before the primary function that handles the endpoint
func Handler (r *chi.Mux) {
	//global middleware(applies to all endpoints)
	r.Use(chimiddle.StripSlashes)
	//set up a route
	r.Route("/account",func(router chi.Router) {
		// middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins",GetCoinBalance)
	})
}