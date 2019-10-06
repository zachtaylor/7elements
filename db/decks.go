package db

import (
	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/db"
)

func NewDeckService(db *db.DB) vii.DeckService {
	return &DeckService{
		conn: db,
	}
}

type DeckService struct {
	conn  *db.DB
	cache vii.Decks
}

func (ds *DeckService) Start() error {
	if decks, err := ds.reloadDecks(); err != nil {
		return err
	} else if deckscards, err := ds.reloadDecksCards(); err != nil {
		return err
	} else {
		for deckid, cards := range deckscards {
			if deck := decks[deckid]; deck != nil {
				deck.Cards = cards
			}
		}
		ds.cache = decks
	}
	return nil
}

func (ds *DeckService) reloadDecks() (vii.Decks, error) {
	decks := make(vii.Decks)
	rows, err := ds.conn.Query("SELECT id, name, cover FROM decks")
	if err != nil {
		return decks, err
	}
	defer rows.Close()
	for rows.Next() {
		deck := &vii.Deck{}
		if err = rows.Scan(&deck.ID, &deck.Name, &deck.CoverID); err == nil {
			decks[deck.ID] = deck
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

func (ds *DeckService) GetAll() (vii.Decks, error) {
	if ds.cache == nil {
		ds.Start()
	}
	return ds.cache, nil
}

func (ds *DeckService) Get(id int) (*vii.Deck, error) {
	decks, err := ds.GetAll()
	if decks == nil {
		return nil, err
	}
	return decks[id], err
}
