package game

import (
	"time"

	"github.com/zachtaylor/7elements/account"
)

type Search struct {
	Deck     *account.Deck
	Start    time.Time
	Done     chan string
	Settings SearchSettings
}

type SearchSettings struct {
	UseP2P bool
}

func NewSearch(deck *account.Deck) *Search {
	return &Search{
		Deck:     deck,
		Start:    time.Now(),
		Done:     make(chan string),
		Settings: SearchSettings{},
	}
}
