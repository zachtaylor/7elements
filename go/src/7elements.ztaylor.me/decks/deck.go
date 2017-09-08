package decks

import (
	"time"
	"ztaylor.me/json"
)

type Deck struct {
	Name     string
	Username string
	Id       int
	Register time.Time
	Cards    map[int]int
	Wins     int
}

func New() *Deck {
	return &Deck{
		Cards: make(map[int]int),
	}
}

func (deck *Deck) Count() int {
	total := 0
	for _, count := range deck.Cards {
		total += count
	}
	return total
}

func (deck *Deck) Json() json.Json {
	return json.Json{
		"id":       deck.Id,
		"username": deck.Username,
		"name":     deck.Name,
		"cards":    deck.Cards,
		"wins":     deck.Wins,
	}
}
