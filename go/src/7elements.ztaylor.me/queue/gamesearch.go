package queue

import (
	"7elements.ztaylor.me/decks"
	"time"
	"ztaylor.me/http/sessions"
)

type GameSearch struct {
	*sessions.Session
	*decks.Deck
	Start time.Time
	Done  chan int
}

func NewGameSearch(session *sessions.Session, deck *decks.Deck) *GameSearch {
	return &GameSearch{
		Session: session,
		Deck:    deck,
		Start:   time.Now(),
		Done:    make(chan int),
	}
}
