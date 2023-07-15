package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func RegisterRoutes(router *mux.Router) {
		router.HandleFunc("/player", controllers.GetAllUsers).Methods("GET")
}