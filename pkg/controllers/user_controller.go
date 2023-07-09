package controllers

import (
	"encoding/json"
	"net/http"
	"football_server/pkg/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Retrieve users from the database or any other data source
	users := []models.User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}

	// Convert users to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}