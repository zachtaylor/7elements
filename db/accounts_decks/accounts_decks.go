package accounts_decks

import (
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
	"ztaylor.me/db"
)

func Get(conn *db.DB, username string) (deck.Prototypes, error) {
	decks := make(deck.Prototypes)

	// decks
	rows, err := conn.Query(
		"SELECT id, name, cover FROM accounts_decks WHERE username=?",
		username,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		deck := deck.NewPrototype()
		err = rows.Scan(&deck.ID, &deck.Name, &deck.Cover)
		if err != nil {
			return nil, err
		}
		decks[deck.ID] = deck
	}
	rows.Close()

	// decks_items
	rows, err = conn.Query("SELECT id, cardid, amount FROM accounts_decks_items WHERE deckid IN (SELECT id FROM decks WHERE username=?)",
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
