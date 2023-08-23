package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Player struct {
	PlayerId int `json:"player_id"`
	Player_name string `json:"player_name"`
    Player_team string `json:"team_name"`	

};
type Team struct {
	TeamId int `json:"team_id"`
	TeamName string `json:"team_name"`
	Abbreviation string `json:"abbreviation"`
	HexCode string `json:"hex_code"`
	Logo string `json:"logo"`
	Stadium string `json:"stadium"`
}

type Stats struct {
	PlayerName   string `json:"player_name"`
	Games        int    `json:"games"`
	MinsPlayed   int    `json:"mins_played"`
	Goal         int    `json:"goal"`
	Assists      int    `json:"assists"`
	Shots        int    `json:"shots"`
	KeyPasses    int    `json:"key_passes"`
	YellowCards  int    `json:"yellow_cards"`
	RedCards     int    `json:"red_cards"`
	Position     string `json:"position"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "golyanglyang"
	dbname   = "try"
)

func main() {
	http.HandleFunc("/teams/", retrieveTeams)	
	http.HandleFunc("/players/", retrievePlayers)	
	http.HandleFunc("/stats/", retrieveStats)	
	
	

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
	query := fmt.Sprintf("SELECT team_id, team_name, abbreviation, hex_code,logo, stadium from teams;")
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	teams := []Team{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var team Team
		err := rows.Scan(&team.TeamId, &team.TeamName, &team.Abbreviation, &team.HexCode, &team.Logo, &team.Stadium)
		if err != nil {
			log.Println(err)
			continue
		}
		teams = append(teams, team)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
func retrievePlayers(w http.ResponseWriter, r *http.Request) {
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
	query := fmt.Sprintf("SELECT player_id, player_name, team_name FROM player ORDER by team_id ASC;")
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
		err := rows.Scan(&player.PlayerId,&player.Player_name, &player.Player_team)
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



func retrieveStats(w http.ResponseWriter, r *http.Request) {
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
	query := fmt.Sprintf("SELECT p.player_name, s.games ,s.Mins_played, s.goal , s.assists, s.shots, s.key_passes , s.yellow_cards, s.red_cards, s.position FROM player p JOIN stats s On p.player_id=s.player_id;")
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Stats{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Stats
		err := rows.Scan(&player.PlayerName, &player.Games, &player.MinsPlayed, &player.Goal, &player.Assists, &player.Shots, &player.KeyPasses, &player.YellowCards, &player.RedCards, &player.Position)
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