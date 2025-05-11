package postgres

import (
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/jmoiron/sqlx"
)

type NHLList struct {
	db *sqlx.DB
}

func NewNHLList(db *sqlx.DB) *NHLList {
	return &NHLList{db: db}
}

// GetTeams take all NHL teams from db
func (n *NHLList) GetTeams() ([]model.NHLTeamsOutput, error) {
	var teams []model.NHLTeamsOutput

	// TODO: fix it
	q := `SELECT t.img_url, t.abbreviation, t.name, n.conference, n.division
      FROM teams t 
      JOIN nhl_teams n ON t.id = n.id_team`

	err := n.db.Select(&teams, q)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// GetSchedule it's util for admin panel
func (n *NHLList) GetSchedule() ([]model.NHLScheduleOutput, error) {
	var nhlGames []model.NHLScheduleOutput

	q := `SELECT
		ns.date_game,
		ns.time_game,
		vt.name AS visitor_team,
		ht.name AS home_team,
		ns.visitor_score,
		ns.home_score,
		ns.attendance,
		ns.game_duration,
		ns.is_overtime,
		ns.venue,
		ns.notes
		FROM nhl_schedule ns
		JOIN
			nhl_teams nvt ON nvt.id = ns.visitor_team_id
		JOIN
			teams vt ON vt.id = nvt.id_team
		JOIN
			nhl_teams nht ON nht.id = ns.home_team_id
		JOIN
			teams ht ON ht.id = nht.id_team
			ORDER BY ns.date_game DESC`

	err := n.db.Select(&nhlGames, q)
	if err != nil {
		return nil, err
	}

	return nhlGames, nil
}

// GetLastSchedule TODO: refactor it
func (n *NHLList) GetLastSchedule(count int) ([]model.NHLScheduleOutput, error) {
	var nhlGames []model.NHLScheduleOutput

	q := `SELECT
		ns.date_game,
		ns.time_game,
		vt.name AS visitor_team,
		ht.name AS home_team,
		ns.visitor_score,
		ns.home_score,
		ns.attendance,
		ns.game_duration,
		ns.is_overtime,
		ns.venue,
		ns.notes
		FROM nhl_schedule ns
		JOIN
			nhl_teams nvt ON nvt.id = ns.visitor_team_id
		JOIN
			teams vt ON vt.id = nvt.id_team
		JOIN
			nhl_teams nht ON nht.id = ns.home_team_id
		JOIN
			teams ht ON ht.id = nht.id_team
			ORDER BY ns.date_game DESC
		LIMIT $1`

	err := n.db.Select(&nhlGames, q, count)
	if err != nil {
		return nil, err
	}

	return nhlGames, nil
}
