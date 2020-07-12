package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/l33m4n123/pet24/api/v1/handler"
	"github.com/l33m4n123/pet24/api/v1/middlewares"
)

// SetupRoutes initiates all the routes at startup
func SetupRoutes(r *mux.Router) {
	log.Println("Setting up routes")

	r.HandleFunc("/api/v1/shelters", handler.ShelterHandler)
}

// SetupMiddlewares initiates all Middlewares for the router
func SetupMiddlewares(r *mux.Router) {
	log.Println("Setting up Middlewares")

	r.Use(middlewares.JSONMw)
	r.Use(middlewares.LoggingMw)
}
