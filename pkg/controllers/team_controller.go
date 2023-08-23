package controllers

import (
	"encoding/json"
	"net/http"
	"football_server/pkg/models"
	// "database/sql"
	"log"
	"fmt"
	"strings"
	
)
func RetrieveTeams(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		// Query the database
		rows, err := DB.Query(`SELECT * FROM teams;`)
		fmt.Println(rows)
	if err != nil {
		fmt.Println("error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	teams := []models.Team{}
	fmt.Println("hellworold")

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var TeamId models.Team
		err := rows.Scan(&TeamId.TeamId, &TeamId.TeamName, &TeamId.Abbreviation, &TeamId.HexCode, &TeamId.Logo, &TeamId.Stadium)
		if err != nil {
			log.Println(err)
			continue
		}
		teams = append(teams, TeamId)
	}
	fmt.Println("helloworld 2")

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

func RetrieveTeamPlayer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/team/")
	fmt.Println(id)
	// Check if the request is a "GET" request
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	// Query the database
	query := fmt.Sprintf(`SELECT p.player_id, p.player_name,t.hex_code, s.games, s.mins_played, s.goal, s.assists, s.shots, s.key_passes, s.yellow_cards, s.red_cards, s.position
	FROM player AS p
	JOIN stats AS s ON p.player_id = s.player_id
	JOIN teams AS t ON p.team_id = t.team_id
	WHERE p.team_id = '%s';`, id) 
	rows, err := DB.Query(query)
	fmt.Println(rows)
if err != nil {
	fmt.Println("error")
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
defer rows.Close()

// Create a slice to store the retrieved records
teams := []models.TeamPlayerStats{}
fmt.Println("hellworold")

// Iterate over the rows and populate the slice
for rows.Next() {
	var PlayerStats models.TeamPlayerStats
	err := rows.Scan(&PlayerStats.PlayerId, &PlayerStats.PlayerName, &PlayerStats.HexCode, &PlayerStats.Games, &PlayerStats.MinsPlayed, &PlayerStats.Goal, &PlayerStats.Assists, &PlayerStats.Shots, &PlayerStats.KeyPasses, &PlayerStats.YellowCards, &PlayerStats.RedCards, &PlayerStats.Position)
	if err != nil {
		log.Println(err)
		continue
	}
	teams = append(teams, PlayerStats)
}
fmt.Println("helloworld 2")

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