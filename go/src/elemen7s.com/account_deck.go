package vii

import (
	"strconv"
	"time"
	"ztaylor.me/js"
)

type AccountDeck struct {
	Id       int
	Version  string
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

func (deck *AccountDeck) Json() js.Object {
	return js.Object{
		"id":       deck.Id,
		"version":  deck.Version,
		"name":     deck.Name,
		"username": deck.Username,
		"cards":    deck.Cards,
		"wins":     deck.Wins,
		"color":    deck.Color,
	}
}

type AccountDecks []*AccountDeck

func (decks AccountDecks) Json() js.Object {
	data := js.Object{}
	for _, deck := range decks {
		data[strconv.Itoa(deck.Id)] = deck.Json()
	}
	return data
}

var AccountDeckService interface {
	Get(username string) (AccountDecks, error)
	Forget(username string)
	Update(deck *AccountDeck) error
	UpdateName(username string, id int, name string) error
	UpdateTallyWinCount(username string, deckid int, version string) error
}
