package queue

import (
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/games"
	"github.com/cznic/mathutil"
	"sync"
	"time"
	"ztaylor.me/events"
	"ztaylor.me/http/sessions"
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
var gameIdGen, _ = mathutil.NewFC32(0, 999999999, true)

func Start(session *sessions.Session, deck *decks.Deck) *GameSearch {
	qlock.Lock()
	defer qlock.Unlock()
	search := NewGameSearch(session, deck)
	queue = append(queue, search)
	log.Add("Username", search.Session.Username).Add("DeckId", deck.Id).Info("queue: start")
	return search
}

func Remove(session *sessions.Session) {
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

	for i := 1; len(queue) > 1; i++ {
		if s1, s2, qlen := queue[0], queue[i], len(queue); TestMatch(queue[0], queue[i]) {
			i = 0
			queue[0] = queue[qlen-1]
			queue[i] = queue[qlen-2]
			queue[qlen-1] = nil
			queue[qlen-2] = nil
			queue = queue[:qlen-2]
			match(s1, s2)
		}
	}
}

func TestMatch(s1 *GameSearch, s2 *GameSearch) bool {
	return true
}

func match(s1 *GameSearch, s2 *GameSearch) {
	game := games.New()
	game.Id = int(gameIdGen.Next())
	games.GetActiveGames(s1.Session.Username)[game.Id] = true
	games.GetActiveGames(s2.Session.Username)[game.Id] = true

	game.Seats = append(game.Seats, games.BuildGameSeat(s1.Deck, game))
	game.Seats = append(game.Seats, games.BuildGameSeat(s2.Deck, game))

	games.Cache[game.Id] = game
	s1.Done <- game.Id
	close(s1.Done)
	s2.Done <- game.Id
	go close(s2.Done)
	events.Fire("GameStart", game)
}
