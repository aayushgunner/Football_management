package main
 


import (
    "database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	_ "github.com/lib/pq"
)
type Player struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "Helloworld"
    dbname   = "try"
)

func retrieveRecord(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         // close database

		// Open the database
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer db.Close()
	
		// Check the database connection
		err = db.Ping()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	
		// Query the database
		rows, err := db.Query(`SELECT "first_name", "last_name" FROM "players"`)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
	
		// Create an array to hold the players
		players := []Player{}
	
		// Loop through the rows and populate the player array
		for rows.Next() {
			player := Player{}
			err := rows.Scan(&player.FirstName, &player.LastName)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			players = append(players, player)
		}
	
		// Check for any errors during the rows iteration
		if err = rows.Err(); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	
		// Marshal the players array to JSON
		jsonData, err := json.Marshal(players)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
	
		// Write the JSON response
		w.Write(jsonData)
	}
func main() {
	fmt.Println("This is the main function.")

	http.HandleFunc("/", retrieveRecord)
	http.ListenAndServe(":8000", nil)
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}