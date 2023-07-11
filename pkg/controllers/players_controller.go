package controllers

import (
	"encoding/json"
	"net/http"
	"football_server/pkg/models"
	// "database/sql"
	"log"
	"fmt"
	
)
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			// fmt.Println("hellworold")

			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		// Query the database
		rows, err := DB.Query(`SELECT "first_name","last_name", "player_id","position","team_id" FROM "players"`)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create an array to hold the players
		players := []models.Player{}

		// Loop through the rows and populate the player array
		for rows.Next() {
			player := models.Player{}
			err := rows.Scan(&player.FirstName, &player.LastName, &player.PlayerId, &player.Position, &player.TeamId)
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