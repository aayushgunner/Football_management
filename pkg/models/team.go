package models
type Team struct {
	TeamId int `json:"team_id"`
	TeamName string `json:"Team_name"`
	Abbreviation string `json:"Abbreviation"`
	HexCode string `json:"hex_code"`
	Logo string `json:"logo"`
	Stadium string `json:"Stadium"`
}