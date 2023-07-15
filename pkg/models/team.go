package models
type Team struct {
	TeamId int `json:"team_id"`
	TeamName string `json:"team_name"`
	Abbreviation string `json:"abbreviation"`
	HexCode string `json:"hex_code"`
	Logo string `json:"logo"`
	Stadium string `json:"stadium"`
}