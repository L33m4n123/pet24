package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/l33m4n123/pet24/api/handler"
	"github.com/l33m4n123/pet24/api/middlewares"
)

// SetupRoutes initiates all the routes at startup
func SetupRoutes(r *mux.Router) {
	log.Println("Setting up routes")

	r.Use(middlewares.JSONMw)
	r.Use(middlewares.LoggingMw)
	r.HandleFunc("/api/v1/shelters", handler.ShelterHandler)
}
