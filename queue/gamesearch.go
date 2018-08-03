package queue

import (
	"github.com/zachtaylor/7tcg"
	"time"
	"ztaylor.me/http"
)

type GameSearch struct {
	*http.Session
	Deck  *vii.AccountDeck
	Start time.Time
	Done  chan string
}

func NewGameSearch(session *http.Session, deck *vii.AccountDeck) *GameSearch {
	qlock.Lock()
	defer qlock.Unlock()

	if HasSearch(session.Username) {
		return nil
	}

	return &GameSearch{
		Session: session,
		Deck:    deck,
		Start:   time.Now(),
		Done:    make(chan string),
	}
}
