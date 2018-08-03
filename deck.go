package vii

import "strconv"

// Deck is a system premade deck
type Deck struct {
	ID    int
	Name  string
	Wins  int
	Color string
	Cards map[int]int
}

func NewDeck() *Deck {
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

func (deck *Deck) Json() Json {
	return Json{
		"id":    deck.ID,
		"name":  deck.Name,
		"wins":  deck.Wins,
		"color": deck.Color,
		"cards": deck.Cards,
	}
}

type Decks map[int]*Deck

func (decks Decks) Json() Json {
	data := Json{}
	for _, deck := range decks {
		data[strconv.Itoa(deck.ID)] = deck.Json()
	}
	return data
}

// DeckService provides access to Decks
var DeckService interface {
	GetAll() (Decks, error)
	Get(int) (*Deck, error)
}
