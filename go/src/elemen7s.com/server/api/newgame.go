package api

import (
	"elemen7s.com/accounts"
	"elemen7s.com/decks"
	"elemen7s.com/games"
	"elemen7s.com/queue"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func NewGameHandler(r *http.Request) error {
	var gameid int
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := accounts.Get(r.Username); account == nil {
		return err
	} else if mydecks, err := decks.Get(r.Username); mydecks == nil {
		log.Add("Error", err).Error("newgame: decks missing")
		return err
	} else if deck := mydecks[r.Data.Ival("deckid")]; deck == nil {
		log.Warn("queue: start failed, deck missing")
	} else if deck.Count() < 21 {
		r.Agent.WriteJson(js.Object{
			"uri": "/notification",
			"data": js.Object{
				"class":   "error",
				"title":   "New Game",
				"message": "Your deck must have at least 21 cards",
				"timeout": 7000,
			},
		})
		log.Add("Count", deck.Count()).Warn("queue: start failed, deck too small")
	} else if ai := r.Data.Bval("ai"); ai {
		game := games.BuildAIGame()
		game.Register(deck, "en-US")
		games.Start(game)
		gameid = game.Id
	} else if v, ok := <-queue.Start(r.Session, deck); !ok {
		log.WithFields(log.Fields{
			"Session": r.Session,
			"DeckId":  deck.Id,
		}).Warn("newgame: failed")
	} else {
		gameid = v
		log.Add("GameId", gameid).Info("newgame.json")
	}

	if game := games.Cache.Get(gameid); game != nil {
		r.Agent.WriteJson(js.Object{
			"uri":  "/match",
			"data": game.StateJson(r.Session.Username),
		})
	}

	return nil
}
