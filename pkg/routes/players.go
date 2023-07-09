package routes

import (
	"github.com/gorilla/mux"
	// "football_server/pkg/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	// router.HandleFunc("/about", controllers.AboutHandler).Methods("GET")

	return router
}