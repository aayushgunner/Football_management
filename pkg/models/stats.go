package models
type Stats struct {
	PlayerName   string `json:"player_name"`
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