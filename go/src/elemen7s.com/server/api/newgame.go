package api

import (
	"elemen7s.com"
	"elemen7s.com/ai"
	"elemen7s.com/engine"
	"elemen7s.com/queue"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func NewGameHandler(r *http.Request) error {
	gameid := ""
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
		return err
	} else if mydecks, err := vii.AccountDeckService.Get(r.Username); mydecks == nil {
		log.Add("Error", err).Error("newgame: decks missing")
		return err
	} else if deck := mydecks[r.Data.Ival("deckid")]; deck == nil {
		log.Add("Deck", r.Data.Ival("deckid")).Error("queue: start failed, deck missing")
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
	} else if r.Data.Bval("ai") {
		game := vii.GameService.New()
		gameid = game.Key
		game.Register(deck, "en-US")
		ai.Register(game)
		go engine.Run(game)
		log.WithFields(log.Fields{
			"GameId":   game,
			"Username": r.Username,
		}).Info("newgame.json: created game vs ai")
	} else if v, ok := <-queue.Start(r.Session, deck); !ok {
		log.WithFields(log.Fields{
			"Session": r.Session,
			"DeckId":  deck.Id,
		}).Warn("newgame: failed")
	} else {
		gameid = v
		log.WithFields(log.Fields{
			"GameId":   gameid,
			"Username": r.Username,
		}).Info("newgame.json: found game vs human")
	}

	if gameid == "" {
		return nil
	}
	game := vii.GameService.Get(gameid)
	if game == nil {
		return ErrGameMissing
	}

	r.Agent.WriteJson(js.Object{
		"uri":  "/match",
		"data": game.Json(r.Session.Username),
	})
	return nil
}
