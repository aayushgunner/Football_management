package models
type Table struct {
	TeamRank        int    `json:"team_rank"`
	TeamName      	string   `json:"team_name"`
	Played          int    `json:"matches_played"`
	Wins        	int    `json:"wins"`
	Draws    		int    `json:"draws"`
	Losses  		int    `json:"losses"`
	GoalsFor        int    `json:"goals_for"`
	GoalsAgainst    int    `json:"goals_against"`
	Points   		int    `json:"points"`
	TeamLogo        string `json:"team_logo"`
 }