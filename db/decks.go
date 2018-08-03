package db

import (
	"github.com/zachtaylor/7tcg"
)

func init() {
	vii.DeckService = DeckService{}
}

type DeckService struct {
	vii.Decks
}

func (service DeckService) Start() error {
	decks, err := service.reloadDecks()
	if err != nil {
		return err
	}
	deckscards, err := service.reloadDecksCards()
	if err != nil {
		return err
	}

	for deckid, cards := range deckscards {
		if deck := decks[deckid]; deck != nil {
			deck.Cards = cards
		}
	}

	service.Decks = decks
	return nil
}

func (service DeckService) reloadDecks() (vii.Decks, error) {
	decks := make(vii.Decks)
	rows, err := conn.Query("SELECT id, name, wins, color FROM decks")
	if err != nil {
		return decks, err
	}
	defer rows.Close()
	for rows.Next() {
		deck := &vii.Deck{}
		if err = rows.Scan(&deck.ID, &deck.Name, &deck.Wins, &deck.Color); err == nil {
			decks[deck.ID] = deck
		} else {
			break
		}
	}
	return decks, err
}
func (service DeckService) reloadDecksCards() (map[int]map[int]int, error) {
	deckscards := make(map[int]map[int]int)
	rows, err := conn.Query("SELECT deckid, cardid, amount FROM decks_items")
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

func (service DeckService) GetAll() (vii.Decks, error) {
	if service.Decks == nil {
		service.Start()
	}
	return service.Decks, nil
}

func (service DeckService) Get(id int) (*vii.Deck, error) {
	decks, err := service.GetAll()
	if decks == nil {
		return nil, err
	}
	return decks[id], err
}
