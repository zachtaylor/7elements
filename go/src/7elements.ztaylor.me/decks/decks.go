package decks

import (
	"strconv"
	"ztaylor.me/json"
)

type Decks map[int]*Deck

func NewDecks() Decks {
	return Decks{
		1: &Deck{Id: 1, Cards: make(map[int]int)},
		2: &Deck{Id: 2, Cards: make(map[int]int)},
		3: &Deck{Id: 3, Cards: make(map[int]int)},
	}
}

func (decks Decks) Json() json.Json {
	data := json.Json{}
	for deckid, deck := range decks {
		data[strconv.FormatInt(int64(deckid), 10)] = deck.Json()
	}
	return data
}
