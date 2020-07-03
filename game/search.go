package game

import (
	"time"

	"github.com/zachtaylor/7elements/deck"
)

type Search struct {
	Deck  *deck.T
	Start time.Time
	Done  chan string
}

func NewSearch(deck *deck.T) *Search {
	return &Search{
		Deck:  deck,
		Start: time.Now(),
		Done:  make(chan string),
	}
}
