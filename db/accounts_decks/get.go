package accounts_decks

import (
	"errors"
	"strconv"

	"github.com/zachtaylor/7elements/deck"
	"taylz.io/db"
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
		deck := deck.NewPrototype(username)
		err = rows.Scan(&deck.ID, &deck.Name, &deck.Cover)
		if err != nil {
			return nil, err
		}
		decks[deck.ID] = deck
	}
	rows.Close()

	// decks_items
	rows, err = conn.Query("SELECT id, cardid, amount FROM accounts_decks_items WHERE username=?",
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

		deck := decks[deckid]
		if deck == nil {
			return nil, errors.New("deckid missing:" + strconv.FormatInt(int64(deckid), 10))
		}
		deck.Cards[cardid] = amount
	}
	rows.Close()

	return decks, nil
}
