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

type HomeFixtures struct {
	TeamName    string    `json:"team_home_name"`
	Possession  string   `json:"possession"`
	ShotsOnGoal    int    `json:"Shots_on_goal"`
	TotalShots  int   `json:"total_shots"`
	TotalPasses    int    `json:"total_passes"`
	CornerKicks  string    `json:"corner_kicks"`
	Offsides     int    `json:"offsides"`
	YellowCards     int    `json:"yellow_cards"`
	RedCards     int    `json:"red_cards"`
	Fouls    int    `json:"fouls"`
}
type AwayFixtures struct {
	TeamName     string    `json:"team_away_name"`
	Possession  string   `json:"possession"`
	ShotsOnGoal    int    `json:"Shots_on_goal"`
	TotalShots  int   `json:"total_shots"`
	TotalPasses    int    `json:"total_passes"`
	CornerKicks  string    `json:"corner_kicks"`
	Offsides     int    `json:"offsides"`
	YellowCards     int    `json:"yellow_cards"`
	RedCards     int    `json:"red_cards"`
	Fouls    int    `json:"fouls"`
}