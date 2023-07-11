package models

type Player struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PlayerId string  `json:"player_id"`
	Position string  `json:"position"`
    TeamId string  `json:"team_id"`
}