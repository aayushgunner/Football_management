package models
type Table struct {
	TeamRank        int    `json:"team_rank"`
	Points   		int    `json:"points"`
	GoalDifference  int    `json:"goalDiff"`
	TeamName      	int    `json:"team_name"`
	Wins        	int    `json:"wins"`
	Draws    		int    `json:"draws"`
	Losses  		int    `json:"losses"`
}