package queue

import (
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/server/sessionman"
	"time"
)

type GameSearch struct {
	*sessionman.Session
	*decks.Deck
	Start time.Time
	Done  chan int
}

func NewGameSearch(session *sessionman.Session, deck *decks.Deck) *GameSearch {
	return &GameSearch{
		Session: session,
		Deck:    deck,
		Start:   time.Now(),
		Done:    make(chan int),
	}
}
