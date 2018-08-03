package vii

import (
	"strconv"
	"time"
)

type AccountDeck struct {
	ID       int
	Name     string
	Username string
	Register time.Time
	Cards    map[int]int
	Wins     int
	Color    string
}

func NewAccountDeck() *AccountDeck {
	return &AccountDeck{
		Cards: make(map[int]int),
	}
}

func (deck *AccountDeck) Count() int {
	total := 0
	for _, count := range deck.Cards {
		total += count
	}
	return total
}

func (deck *AccountDeck) Json() Json {
	return Json{
		"id":       deck.ID,
		"name":     deck.Name,
		"username": deck.Username,
		"cards":    deck.Cards,
		"wins":     deck.Wins,
		"color":    deck.Color,
	}
}

type AccountDecks []*AccountDeck

func (decks AccountDecks) Json() Json {
	data := Json{}
	for _, deck := range decks {
		data[strconv.Itoa(deck.ID)] = deck.Json()
	}
	return data
}

var AccountDeckService interface {
	Get(username string) (AccountDecks, error)
	Forget(username string)
	Update(deck *AccountDeck) error
	UpdateName(username string, id int, name string) error
	UpdateTallyWinCount(username string, deckid int) error
}
