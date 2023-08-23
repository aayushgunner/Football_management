package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func StandingsRoutes(router *mux.Router) {
		router.HandleFunc("/standings", controllers.RetrieveTable).Methods("GET")
}