package games

import (
	"time"
	"ztaylor.me/ctxpert"
	"ztaylor.me/log"
)

func RespondTimeout(game *Game) {
	log := log.Add("GameId", game.Id).Add("GamePhase", game.GamePhase)

	if game.GamePhase != GPHSrespond {
		log.Warn("respondtimeout: error")
		return
	}

	for _, seat := range game.Seats {
		if game.Context.Get("respond:"+seat.Username) == nil {
			log.Add("Seat", seat.Username).Warn("respondtimeout: missing response")
			game.Context.Store("respond:"+seat.Username, "timeout")
		}
	}

	doRespondTimeout(game)
	log.Info("respondtimeout")
}

func doRespondTimeout(game *Game) {
	game.GamePhase = GPHSplay
	game.Context = ctxpert.WithNewTimeout(ctxpert.New(), 30*time.Second)
	game.Context.Done(func(ctx *ctxpert.Context) {
		go AttackStart(game)
	})

	for _, seat := range game.Seats {
		if game.Context.Get("respond:"+seat.Username) != nil {
			log.Add("ClockVal", game.Context.Get("respond:"+seat.Username)).Warn("respondtimeout: " + seat.Username + " has value")
		}
	}

	SendAllSeats(game, "turn", MakeTurnJson(game, game.CurrentTurn()))
}
