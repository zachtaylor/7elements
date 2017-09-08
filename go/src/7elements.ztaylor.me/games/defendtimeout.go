package games

import (
	"time"
	"ztaylor.me/ctxpert"
	"ztaylor.me/log"
)

func DefendTimeout(game *Game) {
	log := log.Add("GameId", game.Id).Add("GamePhase", game.GamePhase)

	if game.GamePhase != GPHSplay {
		log.Warn("defendtimeout: error")
		return
	} else if game.TurnPhase != TPHSdefend {
		log.Add("TurnPhase", game.TurnPhase).Warn("defendtimeout: error")
		return
	}

	game.TurnPhase = TPHSdone
	game.Context = ctxpert.WithNewTimeout(ctxpert.New(), 4*time.Second)
	game.Context.Always(func(ctx *ctxpert.Context) {
		go DoneTimeout(game)
	})

	turndata := MakeTurnJson(game, game.CurrentTurn())

	for _, seat := range game.Seats {
		if game.Context.Get("hand:"+seat.Username) != nil {
			log.Clone().Add("ClockVal", game.Context.Get("defend:"+seat.Username)).Warn("defendtimeout: " + seat.Username + " already completed defend")
			continue
		}

		if socket := seat.Socket; socket == nil {
			log.Add("Username", seat.Username).Error("defendtimeout: socket not found")
		} else {
			log.Add("GameId", game.Id).Add("Seat", seat.Username).Warn("defendtimeout: missing response")
			go socket.Send("turn", turndata)
		}
	}

	log.Add("GameId", game.Id).Info("defendtimeout")
}
