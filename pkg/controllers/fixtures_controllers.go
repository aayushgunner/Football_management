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
func RetrieveFixtures(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
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

func RetrieveFixturesById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hellworld")
	id := strings.TrimPrefix(r.URL.Path, "/fixtures/")
	fmt.Println(id)
	
	
		// Check if the request is a "GET" request
		if r.Method != "GET" {
			// fmt.Println("hellworold")

			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		// Query the database
		 	
		query := fmt.Sprintf("SELECT team_home_name, possession, Shots_on_goal, total_shots, total_passes, corner_kicks, offsides, yellow_cards, red_cards, fouls FROM homeFixtures WHERE fixture_id = '%s'", id)
		rows, err := DB.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
	
		// Create a slice to store the retrieved records
		homeFixtures := []models.HomeFixtures{}
	
		// Iterate over the rows and populate the slice
		for rows.Next() {
			var stat models.HomeFixtures
			err := rows.Scan(&stat.TeamName, &stat.Possession, &stat.ShotsOnGoal, &stat.TotalShots, &stat.TotalPasses, &stat.CornerKicks, &stat.Offsides, &stat.YellowCards, &stat.RedCards, &stat.Fouls)
			if err != nil {
				log.Println(err)
				continue
			}
			homeFixtures = append(homeFixtures, stat)
		}
	
		// Check for any errors encountered during row iteration
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
	
		// Check if any stats were retrieved
		if len(homeFixtures) == 0 {
			http.NotFound(w, r)
			return
		}

		// Fetch data from awayFixtures table
		query1 := fmt.Sprintf("SELECT team_away_name, possession, Shots_on_goal, total_shots, total_passes, corner_kicks, offsides, yellow_cards, red_cards, fouls FROM awayFixtures WHERE fixture_id = '%s'", id)
		rows1, err := DB.Query(query1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows1.Close()
	
		// Create a slice to store the retrieved records
		awayFixtures := []models.AwayFixtures{}
	
		// Iterate over the rows and populate the slice
		for rows1.Next() {
			var stat models.AwayFixtures
			err := rows1.Scan(&stat.TeamName, &stat.Possession, &stat.ShotsOnGoal, &stat.TotalShots, &stat.TotalPasses, &stat.CornerKicks, &stat.Offsides, &stat.YellowCards, &stat.RedCards, &stat.Fouls)
			if err != nil {
				log.Println(err)
				continue
			}
			awayFixtures = append(awayFixtures, stat)
		}
	
		// Check for any errors encountered during row iteration
		if err := rows1.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
	
		// Check if any stats were retrieved
		if len(awayFixtures) == 0 {
			http.NotFound(w, r)
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		// Combine data into an object
	fixtureData := map[string]interface{}{
		"homeFixtures": homeFixtures,
		"awayFixtures": awayFixtures,
	}

	// Convert the data to JSON
	jsonData, err := json.MarshalIndent(fixtureData, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
	
		// Write the JSON response
		w.Write(jsonData)
}