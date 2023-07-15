package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func TeamRoutes(router *mux.Router) {
	router.HandleFunc("/teams", controllers.RetrieveTeams).Methods("GET")
}