package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"football_server/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}