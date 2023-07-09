package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
}