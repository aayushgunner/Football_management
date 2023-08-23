package models

type Fixtures struct {
	FixtureId     int    `json:"fixture_id"`
	FixtureDate   string   `json:"fixture_date"`
	TeamHomeID    int    `json:"team_home_id"`
	TeamHomeName  string    `json:"team_home_name"`
	TeamAwayID    int    `json:"team_away_id"`
	TeamAwayName  string    `json:"team_away_name"`
	HomeGoals     int    `json:"home_goals"`
	AwayGoals     int    `json:"away_goals"`
 }