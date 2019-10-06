package vii

import (
	"fmt"
	"sort"
	"strings"

	"ztaylor.me/cast"
)

// Deck is a system premade deck
type Deck struct {
	ID      int
	Name    string
	CoverID int
	Cards   map[int]int
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

func (deck *Deck) JSON() cast.JSON {
	cardsJSON := cast.JSON{}
	for k, v := range deck.Cards {
		cardsJSON[cast.StringI(int(k))] = v
	}
	return cast.JSON{
		"id":    deck.ID,
		"name":  deck.Name,
		"cover": "/img/card/" + cast.StringI(deck.CoverID) + ".jpg",
		"cards": cardsJSON,
	}
}

type Decks map[int]*Deck

func (decks Decks) JSON() fmt.Stringer {
	json := make([]string, 0)
	keys := make([]int, len(decks))
	var i int
	for k := range decks {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	for _, k := range keys {
		json = append(json, decks[k].JSON().String())
	}
	return cast.Stringer(`[` + strings.Join(json, ",") + `]`)
}

// DeckService provides access to Decks
type DeckService interface {
	GetAll() (Decks, error)
	Get(int) (*Deck, error)
}
