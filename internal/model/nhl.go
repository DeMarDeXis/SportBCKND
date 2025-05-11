package model

import "time"

type NHLTeamsOutput struct {
	ImgURL     string `json:"img_url" db:"img_url"`
	Abbr       string `json:"abbr" db:"abbreviation"`
	Name       string `json:"name" db:"name"`
	Conference string `json:"conference" db:"conference"`
	Division   string `json:"division" db:"division"`
}

type NHLScheduleOutput struct {
	//ID           int       `db:"id"`
	Date         string    `json:"date_game" db:"date_game"`
	Time         string    `json:"time_game" db:"time_game"`
	VisitorTeam  string    `json:"visitor_team" db:"visitor_team"`
	HomeTeam     string    `json:"home_team" db:"home_team"`
	VisitorScore *int      `json:"visitor_score" db:"visitor_score"`
	HomeScore    *int      `json:"home_score" db:"home_score"`
	Attendance   *int      `json:"attendance" db:"attendance"`
	GameDuration *string   `json:"game_duration" db:"game_duration"`
	IsOvertime   bool      `json:"is_overtime" db:"is_overtime"`
	Venue        *string   `json:"venue" db:"venue"`
	Notes        *string   `json:"notes" db:"notes"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

//TODO: do it
//type NHLRoster struct {
//}

//TODO: UsersTeams
//type UsersTeams struct {

//TODO: maybe Teams and UsersTeams should be in another file
