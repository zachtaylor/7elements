package queue

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"sync"
	"time"
	"ztaylor.me/events"
	"ztaylor.me/http"
	"ztaylor.me/log"
)

func init() {
	go func() {
		for range time.Tick(time.Second) {
			watch()
		}
	}()
}

var queue = make([]*GameSearch, 0)
var qlock sync.Mutex

func Start(session *http.Session, deck *vii.AccountDeck) chan int {
	c := make(chan int)
	go func() {
		if search := NewGameSearch(session, deck); search == nil {
			log.Add("Username", session.Username).Add("DeckId", deck.Id).Warn("queue: rate limit")
		} else {
			queue = append(queue, search)

			if gameid, ok := <-search.Done; !ok {
				log.Add("Username", session.Username).Add("DeckId", deck.Id).Warn("queue: finished empty handed")
			} else {
				log.Add("Username", session.Username).Add("DeckId", deck.Id).Info("queue: finished")
				c <- gameid
			}
		}
		close(c)
	}()
	return c
}

func HasSearch(username string) bool {
	qlock.Lock()
	defer qlock.Unlock()
	for _, search := range queue {
		if search.Session.Username == username {
			return true
		}
	}
	return false
}

func Remove(session *http.Session) {
	qlock.Lock()
	defer qlock.Unlock()

	for i := 0; i < len(queue); i++ {
		if search := queue[i]; search.Session == session {
			log.Add("Username", session.Username).Add("Timer", time.Now().Sub(search.Start).Seconds()).Info("queue: remove")
			close(queue[i].Done)
			queue[i] = queue[len(queue)-1]
			queue[len(queue)-1] = nil
			queue = queue[0 : len(queue)-1]
			return
		}
	}
}

func watch() {
	qlock.Lock()
	defer qlock.Unlock()

	for i := 1; i < len(queue); i++ {
		if s1, s2, qlen := queue[0], queue[i], len(queue); TestMatch(s1, s2) {
			queue[0] = queue[qlen-1]
			queue[i] = queue[qlen-2]
			queue[qlen-1] = nil
			queue[qlen-2] = nil
			queue = queue[:qlen-2]
			match(s1, s2)
			return
		}
	}
}

func TestMatch(s1 *GameSearch, s2 *GameSearch) bool {
	return s1.Session.Username != s2.Session.Username
}

func match(s1 *GameSearch, s2 *GameSearch) {
	game := games.New()

	game.Register(s1.Deck, "en-US") // right here is where to put user language pref
	game.Register(s2.Deck, "en-US") // for multilanguage cards in the same game
	games.Start(game)

	s1.Done <- game.Id
	close(s1.Done)
	s2.Done <- game.Id
	close(s2.Done)
	events.Fire("GameStart", game)
}
