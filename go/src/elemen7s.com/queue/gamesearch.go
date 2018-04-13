package queue

import (
	"elemen7s.com/decks"
	"time"
	"ztaylor.me/http"
)

type GameSearch struct {
	*http.Session
	*decks.Deck
	Start time.Time
	Done  chan int
}

func NewGameSearch(session *http.Session, deck *decks.Deck) *GameSearch {
	qlock.Lock()
	defer qlock.Unlock()

	if HasSearch(session.Username) {
		return nil
	}

	return &GameSearch{
		Session: session,
		Deck:    deck,
		Start:   time.Now(),
		Done:    make(chan int),
	}
}
