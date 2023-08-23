package models
type Team struct {
	TeamId int `json:"team_id"`
	TeamName string `json:"team_name"`
	Abbreviation string `json:"abbreviation"`
	HexCode string `json:"hex_code"`
	Logo string `json:"logo"`
	Stadium string `json:"stadium"`
}

type TeamPlayerStats struct {
	PlayerId int `json:"player_id"`
	PlayerName   string `json:"player_name"`
	HexCode string `json:"hex_code"`
	Games        int    `json:"games"`
	MinsPlayed   int    `json:"mins_played"`
	Goal         int    `json:"goal"`
	Assists      int    `json:"assists"`
	Shots        int    `json:"shots"`
	KeyPasses    int    `json:"key_passes"`
	YellowCards  int    `json:"yellow_cards"`
	RedCards     int    `json:"red_cards"`
	Position     string `json:"position"`
}