package games

import (
	"time"

	"taylz.io/db"
)

func CountWins(conn *db.DB, username string) (int, error) {
	row := conn.QueryRow(
		"SELECT COUNT(*) FROM games WHERE winner=?",
		username,
	)

	var count int64
	if err := row.Scan(&count); err != nil {
		return 0, err
	} else {
		return int(count), nil
	}
}

func Insert(conn *db.DB, winner string, winner_skill int, loser string, loser_skill int) (err error) {
	_, err = conn.Exec(
		"INSERT INTO games (time, winner, winner_skill, loser, loser_skill) VALUES (?, ?, ?, ?, ?)",
		time.Now().Unix(),
		winner,
		winner_skill,
		loser,
		loser_skill,
	)
	return
}
