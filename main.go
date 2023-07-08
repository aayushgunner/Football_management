package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
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
	password = "golyanglyang"
	dbname   = "try"
)

func main() {
	http.HandleFunc("/records/Mikel", retrieveRecordByMikel)
	http.HandleFunc("/records/Pep", retrieveRecordByPep)
	http.HandleFunc("/records/Arsenal", retrieveRecordByArsenal)
	http.HandleFunc("/records/Liverpool", retrieveRecordByLiverpool)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func retrieveRecordByMikel(w http.ResponseWriter, r *http.Request) {
	retrieveRecord(w, r, "Mikel")
}

func retrieveRecordByPep(w http.ResponseWriter, r *http.Request) {
	retrieveRecord(w, r, "Pep")
}

func retrieveRecordByArsenal(w http.ResponseWriter , r *http.Request) {
	retrieveRecordteam(w, r, "Arsenal")
}

func retrieveRecordByLiverpool(w http.ResponseWriter , r * http.Request) {
	retrieveRecordteam(w , r , "Liverpool")
}
func retrieveRecordteam(w http.ResponseWriter , r *http.Request, teamName string){
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
	query := fmt.Sprintf("SELECT first_name , last_name FROM players where team_id IN (SELECT team_id FROM team  WHERE team_name='%s') ORDER BY first_name ASC;", teamName)
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


func retrieveRecord(w http.ResponseWriter, r *http.Request, coachName string) {
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
	query := fmt.Sprintf("SELECT first_name, last_name FROM players WHERE team_id IN (SELECT team_id FROM team WHERE coach_id IN (SELECT coach_id FROM coach WHERE first_name = '%s')) ORDER BY first_name ASC;", coachName)
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