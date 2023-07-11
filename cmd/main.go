package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"football_server/pkg/routes"
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	"football_server/pkg/controllers"
)
var db *sql.DB
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "Helloworld"
	dbname   = "try"
)
// func setupDatabase() (*sql.DB, error) {
//     // Replace the connection string with your actual PostgreSQL database details
//     db, err := sql.Open("postgres", "postgres://user:password@host:port/database?sslmode=disable")
//     if err != nil {
//         return nil, err
//     }
//     return db, nil
// }
func main() {
	// Create the connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open the database connection
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	controllers.SetDB(db)

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Apply CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Use the CORS middleware with the router
	router.Use(corsHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}