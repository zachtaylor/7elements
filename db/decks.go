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
	conn  *db.DB
	cache deck.Prototypes
}

func (ds *DeckService) Start() error {
	if decks, err := ds.reloadDecks(); err != nil {
		return err
	} else if deckscards, err := ds.reloadDecksCards(); err != nil {
		return err
	} else {
		for deckid, cards := range deckscards {
			if d := decks[deckid]; d != nil {
				d.Cards = cards
			}
		}
		ds.cache = decks
	}
	return nil
}

func (ds *DeckService) reloadDecks() (deck.Prototypes, error) {
	decks := make(deck.Prototypes)
	rows, err := ds.conn.Query("SELECT id, name, cover FROM decks")
	if err != nil {
		return decks, err
	}
	defer rows.Close()
	for rows.Next() {
		d := &deck.Prototype{}
		if err = rows.Scan(&d.ID, &d.Name, &d.CoverID); err == nil {
			decks[d.ID] = d
		} else {
			break
		}
	}
	return decks, err
}
func (ds *DeckService) reloadDecksCards() (map[int]map[int]int, error) {
	deckscards := make(map[int]map[int]int)
	rows, err := ds.conn.Query("SELECT deckid, cardid, amount FROM decks_items")
	if err != nil {
		return deckscards, err
	}
	defer rows.Close()
	for rows.Next() {
		var deckid, cardid, amount int
		if err = rows.Scan(&deckid, &cardid, &amount); err == nil {
			if deckscards[deckid] == nil {
				deckscards[deckid] = make(map[int]int)
			}
			deckscards[deckid][cardid] = amount
		} else {
			break
		}
	}
	return deckscards, err
}

func (ds *DeckService) GetAll() (deck.Prototypes, error) {
	if ds.cache == nil {
		ds.Start()
	}
	return ds.cache, nil
}

func (ds *DeckService) Get(id int) (*deck.Prototype, error) {
	decks, err := ds.GetAll()
	if decks == nil {
		return nil, err
	}
	return decks[id], err
}
