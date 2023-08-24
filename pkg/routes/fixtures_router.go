package routes

import (
	"github.com/gorilla/mux"
	"football_server/pkg/controllers"
)

func FixturesRoutes(router *mux.Router) {
		router.HandleFunc("/fixtures", controllers.RetrieveFixtures).Methods("GET")
		router.HandleFunc("/fixtures/{id}", controllers.RetrieveFixturesById).Methods("GET")
}