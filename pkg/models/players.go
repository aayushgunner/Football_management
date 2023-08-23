package models

type Player struct {
	PlayerId int `json:"player_id"`
	PlayerName  string `json:"player_name"`
	TeamName string  `json:"team_name"`
    TeamId int  `json:"team_id"`
}