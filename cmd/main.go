package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    _ "github.com/lib/pq"
	_ "github.com/gorilla/mux"
	_ "football_server/pkg/routes"
)

type Player struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
};
type Team struct {
	TeamName string `json:"team_name"`
	TeamId int `json:"team_id"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "golyanglyang"
	dbname   = "try"
)
func main() {

	http.HandleFunc("/teams/",retrieveTeams)
	http.HandleFunc("/records/",retrieveRecordByteam)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

	
	
}

func retrieveTeams (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Database connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	query := fmt.Sprintf("SELECT team_id, team_name from team;")
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Team{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Team
		err := rows.Scan(&player.TeamId, &player.TeamName)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

func retrieveRecordByteam (w http.ResponseWriter, r *http.Request) {
    // vars := mux.Vars(r)
    id := r.URL.Path[len("/records/"):]
    // if !ok {
    //     fmt.Println("id is missing in parameters")
    // }
    fmt.Println(`id := `, id)
	retrieveRecord(w,r, id)
    //call http://localhost:8080/provisions/someId in your browser
    //Output : id := someId
	}


func retrieveRecord(w http.ResponseWriter, r *http.Request, id string) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Database connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	query := fmt.Sprintf("SELECT first_name, last_name FROM players WHERE team_id = '%s'", id)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Player{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.FirstName, &player.LastName)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
