package controllers

import (
	"encoding/json"
	"net/http"
	"football_server/pkg/models"
	// "database/sql"
	"log"
	"fmt"
	
)
func RetrieveTable(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			// fmt.Println("hellworold")

			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		
		// Query the database
		rows, err := DB.Query(`SELECT team_rank, team_name, matches_played, wins, draws, losses , goals_for, goals_against, points FROM standings ORDER BY team_rank;`)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create an array to hold the players
		players := []models.Table{}

		// Loop through the rows and populate the player array
		for rows.Next() {
			player := models.Table{}
			err := rows.Scan(&player.TeamRank, &player.TeamName, &player.Played, &player.Wins, &player.Draws, &player.Losses, &player.GoalsFor, &player.GoalsAgainst, &player.Points)
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