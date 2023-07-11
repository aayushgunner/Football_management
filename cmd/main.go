package main

import (
	"log"
	"net/http"
	"encoding/csv"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"football_server/pkg/routes"
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	"os"
	"io"
	"path/filepath"
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
	absPath, err := filepath.Abs("cmd/clubInfo.csv")
if err != nil {
	log.Fatal(err)
}
file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	controllers.SetDB(db)
	for {
		// Read each row from the CSV file
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	
		fmt.Println("helloworld")
		// Prepare the SQL INSERT statement
		stmt, err := db.Prepare("INSERT INTO team (team_name, abbreviation, hex_code, team_logo, Stadium) VALUES ($1, $2, $3, $4, $5)")
		if err != nil {
			log.Fatal(err)
		}
	
		// Execute the INSERT statement with row data
		fmt.Println(row[0], row[1], row[2], row[3], row[4])
		_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4])
		
		if err != nil {
			log.Fatal(err)
		}
	}

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