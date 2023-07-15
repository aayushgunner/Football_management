package controllers

import (
	"encoding/json"
	"net/http"
	"football_server/pkg/models"
	// "database/sql"
	"log"
	// "fmt"
	
)
func RetrieveTeams(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			// fmt.Println("hellworold")

			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		// Query the database
		rows, err := DB.Query(`SELECT team_id, team_name, abbreviation, hex_code,logo, stadium from teams;`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	teams := []models.Team{}

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