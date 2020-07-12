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

	r.HandleFunc("/shelters", handler.GetShelters).Methods("GET")
	r.HandleFunc("/shelter/{id}", handler.GetShelterByID).Methods("GET")
	r.HandleFunc("/shelters/add", handler.AddNewShelter).Methods("POST")
	r.HandleFunc("/shelters/update/{id}", handler.UpdateShelter).Methods("POST")
	r.HandleFunc("/shelters/delete/{id}", handler.DeleteShelter).Methods("DELETE")
}

// SetupMiddlewares initiates all Middlewares for the router
func SetupMiddlewares(r *mux.Router) {
	log.Println("Setting up Middlewares")

	r.Use(middlewares.JSONMw)
	r.Use(middlewares.LoggingMw)
}
