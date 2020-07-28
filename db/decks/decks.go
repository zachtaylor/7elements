package decks

import (
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/db"
)

// func (ds *DeckService) Start() error {
// 	if decks, err := ds.reloadDecks(); err != nil {
// 		return err
// 	} else if deckscards, err := ds.reloadDecksCards(); err != nil {
// 		return err
// 	} else {
// 		for deckid, cards := range deckscards {
// 			if d := decks[deckid]; d != nil {
// 				d.Cards = cards
// 			}
// 		}
// 		ds.cache = decks
// 	}
// 	return nil
// }

// func (ds *DeckService) reloadDecks() (deck.Prototypes, error) {
// 	decks := make(deck.Prototypes)
// 	rows, err := ds.conn.Query("SELECT id, name, cover FROM decks")
// 	if err != nil {
// 		return decks, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		d := &deck.Prototype{}
// 		if err = rows.Scan(&d.ID, &d.Name, &d.CoverID); err == nil {
// 			decks[d.ID] = d
// 		} else {
// 			break
// 		}
// 	}
// 	return decks, err
// }

func Get(conn *db.DB, id int) (*deck.Prototype, error) {
	row := conn.QueryRow(
		"SELECT id, name, cover FROM decks WHERE id=?",
		id,
	)
	deck := deck.NewPrototype()
	err := row.Scan(&deck.ID, &deck.Name, &deck.Cover)
	if err != nil {
		return nil, err
	}

	// decks_items
	rows, err := conn.Query("SELECT cardid, amount FROM decks_items WHERE deckid=?",
		id,
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

		deck.Cards[cardid] = amount
	}
	rows.Close()

	return deck, nil
}

// func (ds *DeckService) Update(deck *deck.Prototype) (err error) {
// 	if err = ds.Delete(deck.ID); err != nil {
// 	} else {
// 		err = ds.Insert(deck)
// 	}
// 	return
// }

func UpdateName(conn *db.DB, id int, newname string) error {
	_, err := conn.Exec(
		"UPDATE decks SET name=? WHERE id=?",
		newname,
		id,
	)
	return err
}

func Insert(conn *db.DB, deck *deck.Prototype) error {
	_, err := conn.Exec("INSERT INTO decks (id, name, user, cover, wins, loss) VALUES (?, ?, ?, ?, ?, ?)",
		deck.ID,
		deck.Name,
		deck.User,
		deck.Cover,
		deck.Wins,
		deck.Loss,
	)
	if err != nil {
		return err
	}

	for cardId, amount := range deck.Cards {
		_, err := conn.Exec("INSERT INTO decks_items(deckid, cardid, amount) VALUES (?, ?, ?)",
			deck.ID,
			cardId,
			amount,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func Delete(conn *db.DB, deckid int) (err error) {
	_, err = conn.Exec("DELETE FROM decks WHERE AND id=?",
		deckid,
	)
	if err == nil {
		_, err = conn.Exec("DELETE FROM decks_items WHERE id=?",
			deckid,
		)
	}
	return
}

// func (ds *DeckService) reloadDecksCards() (map[int]map[int]int, error) {
// 	deckscards := make(map[int]map[int]int)
// 	rows, err := ds.conn.Query("SELECT deckid, cardid, amount FROM decks_items")
// 	if err != nil {
// 		return deckscards, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var deckid, cardid, amount int
// 		if err = rows.Scan(&deckid, &cardid, &amount); err == nil {
// 			if deckscards[deckid] == nil {
// 				deckscards[deckid] = make(map[int]int)
// 			}
// 			deckscards[deckid][cardid] = amount
// 		} else {
// 			break
// 		}
// 	}
// 	return deckscards, err
// }
