package controllers

import (
	"encoding/json"
	"net/http"
	"football_server/pkg/models"
	// "database/sql"
	"log"
	"fmt"
	
)
func RetrieveFixtures(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			// fmt.Println("hellworold")

			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		
		// Query the database
		rows, err := DB.Query(`SELECT fixture_id, fixture_date, team_home_id, team_home_name, team_away_id, team_away_name, home_goals, away_goals FROM fixtures ORDER BY fixture_id;`)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create an array to hold the players
		players := []models.Fixtures{}

		// Loop through the rows and populate the player array
		for rows.Next() {
			player := models.Fixtures{}
			err := rows.Scan(&player.FixtureId, &player.FixtureDate, &player.TeamHomeID, &player.TeamHomeName, &player.TeamAwayID, &player.TeamAwayName, &player.HomeGoals, &player.AwayGoals)
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