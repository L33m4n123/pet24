package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l33m4n123/pet24/api/v1/config"
	"github.com/l33m4n123/pet24/api/v1/database"
	"github.com/l33m4n123/pet24/api/v1/routes"
)

func main() {
	config.ReadInConfig()
	database.SetupConnection(config.Cfg.Database.Connection)

	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	routes.SetupMiddlewares(apiV1Router)
	routes.SetupRoutes(apiV1Router)

	log.Println("Starting server at http://" + config.Cfg.Server.Host + ":" + config.Cfg.Server.Port)
	http.ListenAndServe(":"+config.Cfg.Server.Port, router)
	database.CloseConnection()
}
