package game

import (
	"time"

	"github.com/zachtaylor/7elements"
)

type Search struct {
	Deck     *vii.AccountDeck
	Start    time.Time
	Done     chan string
	Settings SearchSettings
}

type SearchSettings struct {
	UseP2P bool
}

func NewSearch(deck *vii.AccountDeck) *Search {
	return &Search{
		Deck:     deck,
		Start:    time.Now(),
		Done:     make(chan string),
		Settings: SearchSettings{},
	}
}
