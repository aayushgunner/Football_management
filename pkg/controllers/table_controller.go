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
func RetrieveTable(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			// fmt.Println("hellworold")

			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		// Query the database
		query := fmt.Sprintf("SELECT games, mins_played, goal, assists, shots, key_passes, yellow_cards, red_cards, position FROM stats WHERE player_id = '%s' LIMIT 1;", id)
		rows, err := DB.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
	
		// Create a slice to store the retrieved records
		stats := []models.Stats{}
	
		// Iterate over the rows and populate the slice
		for rows.Next() {
			var stat models.Stats
			err := rows.Scan(&stat.Games, &stat.MinsPlayed, &stat.Goal, &stat.Assists, &stat.Shots, &stat.KeyPasses, &stat.YellowCards, &stat.RedCards, &stat.Position)
			if err != nil {
				log.Println(err)
				continue
			}
			stats = append(stats, stat)
		}
	
		// Check for any errors encountered during row iteration
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
	
		// Check if any stats were retrieved
		if len(stats) == 0 {
			http.NotFound(w, r)
			return
		}
	
		// Convert the data to JSON
		jsonData, err := json.Marshal(stats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
	
		// Write the JSON response
		w.Write(jsonData)
}