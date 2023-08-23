package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func TableRoutes(router *mux.Router) {
		router.HandleFunc("/table", controllers.RetrieveTable).Methods("GET")
}