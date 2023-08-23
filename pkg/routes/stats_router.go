package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func StatsRoutes(router *mux.Router) {
	router.HandleFunc("/stats/{id}", controllers.RetrieveStatsById).Methods("GET")
}