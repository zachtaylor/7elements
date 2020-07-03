package db

import (
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
	"ztaylor.me/db"
)

func getUserDecks(conn *db.DB, username string) (deck.Prototypes, error) {
	decks := make(deck.Prototypes)

	// decks
	rows, err := conn.Query(
		"SELECT id, name, user, cover, wins, loss FROM decks WHERE user=?",
		username,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		deck := deck.NewPrototype()
		err = rows.Scan(&deck.ID, &deck.Name, &deck.User, &deck.Cover, &deck.Wins, &deck.Loss)
		if err != nil {
			return nil, err
		}
		decks[deck.ID] = deck
	}
	rows.Close()

	// decks_items
	rows, err = conn.Query("SELECT deckid, cardid, amount FROM decks_items WHERE deckid IN (SELECT id FROM decks WHERE user=?)",
		username,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var deckid, cardid, amount int

		err = rows.Scan(&deckid, &cardid, &amount)
		if err != nil {
			return nil, err
		}

		if deck := decks[deckid]; deck == nil {
			return nil, cast.NewError(nil, "deckid missing:", deckid)
		} else {
			deck.Cards[cardid] = amount
		}
	}
	rows.Close()

	return decks, nil
}
