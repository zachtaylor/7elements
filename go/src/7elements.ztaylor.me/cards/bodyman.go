package cards

import (
	"7elements.ztaylor.me/db"
)

var BodyCache = make(map[int]*Body)

func LoadBodyCache() error {
	rows, err := db.Connection.Query("SELECT cardid, attack, health FROM cards_bodies")
	if err != nil {
		return err
	}

	for rows.Next() {
		body := &Body{}
		err = rows.Scan(&body.CardId, &body.Attack, &body.Health)
		if err != nil {
			return err
		}

		BodyCache[body.CardId] = body
	}
	rows.Close()

	return nil
}
