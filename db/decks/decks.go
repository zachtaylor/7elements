package decks

import (
	"strconv"

	"github.com/zachtaylor/7elements/deck"
	"taylz.io/db"
	"taylz.io/types"
)

func GetAll(conn *db.DB) (deck.Prototypes, error) {
	rows, err := conn.Query("SELECT id, name, cover FROM decks")
	if err != nil {
		return nil, err
	}
	decks := make(deck.Prototypes)
	for rows.Next() {
		d := deck.NewPrototype("vii")
		if err = rows.Scan(&d.ID, &d.Name, &d.Cover); err == nil {
			decks[d.ID] = d
		} else {
			rows.Close()
			return nil, err
		}
	}
	rows.Close()

	if err = getItems(conn, decks); err != nil {
		return nil, err
	}

	return decks, err
}

func getItems(conn *db.DB, decks deck.Prototypes) error {
	// decks_items
	rows, err := conn.Query("SELECT deckid, cardid, amount FROM decks_items")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var deckid, cardid, amount int

		err = rows.Scan(&deckid, &cardid, &amount)
		if err != nil {
			return err
		}

		deck := decks[deckid]
		if deck == nil {
			return types.NewErr("deckid missing: " + strconv.FormatInt(int64(deckid), 10))
		}

		deck.Cards[cardid] = amount
	}

	return nil
}

// func Get(conn *db.DB, id int) (*deck.Prototype, error) {
// 	row := conn.QueryRow(
// 		"SELECT id, name, cover FROM decks WHERE id=?",
// 		id,
// 	)
// 	deck := deck.NewPrototype()
// 	err := row.Scan(&deck.ID, &deck.Name, &deck.Cover)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return deck, nil
// }

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
