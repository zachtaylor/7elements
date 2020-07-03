package db

import (
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/db"
)

func NewDeckService(db *db.DB) deck.PrototypeService {
	return &DeckService{
		conn: db,
	}
}

type DeckService struct {
	conn *db.DB
}

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

func (ds *DeckService) GetUser(user string) (deck.Prototypes, error) {
	return getUserDecks(ds.conn, user)
}

// func (ds *DeckService) Update(deck *deck.Prototype) (err error) {
// 	if err = ds.Delete(deck.ID); err != nil {
// 	} else {
// 		err = ds.Insert(deck)
// 	}
// 	return
// }

func (ds *DeckService) UpdateName(id, newname string) error {
	_, err := ds.conn.Exec(
		"UPDATE decks SET name=? WHERE id=?",
		newname,
		id,
	)
	return err
}

func (ds *DeckService) Insert(deck *deck.Prototype) error {
	_, err := ds.conn.Exec("INSERT INTO decks (id, name, user, cover, wins, loss) VALUES (?, ?, ?, ?, ?, ?)",
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
		_, err := ds.conn.Exec("INSERT INTO decks_items(deckid, cardid, amount) VALUES (?, ?, ?)",
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

func (ds *DeckService) Delete(deckid string) (err error) {
	_, err = ds.conn.Exec("DELETE FROM decks WHERE AND id=?",
		deckid,
	)
	if err == nil {
		_, err = ds.conn.Exec("DELETE FROM decks_items WHERE id=?",
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
