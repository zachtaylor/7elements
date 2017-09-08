package games

import (
	"ztaylor.me/log"
)

func DoneTimeout(game *Game) {
	log := log.Add("GameId", game.Id)

	if game.GamePhase != GPHSplay {
		log.Add("GamePhase", game.GamePhase).Warn("donetimeout: event rejected")
		return
	}
	if game.TurnPhase != TPHSdone {
		log.Add("TurnPhase", game.TurnPhase).Warn("donetimeout: event rejected")
		return
	}

	turndata := MakeTurnJson(game, game.CurrentTurn())

	for _, seat := range game.Seats {
		if game.Context.Get("hand:"+seat.Username) != nil {
			log.Clone().Add("ClockVal", game.Context.Get("defend:"+seat.Username)).Warn("donetimeout: " + seat.Username + " already completed done")
			continue
		}

		if socket := seat.Socket; socket == nil {
			log.Clone().Add("Username", seat.Username).Error("donetimeout: socket not found")
		} else {
			socket.Send("turn", turndata)
			log.Clone().Add("GameId", game.Id).Add("Seat", seat.Username).Warn("donetimeout: missing response")
		}
	}

	game.TurnPhase = TPHSwait
	log.Add("GameId", game.Id).Info("donetimeout")
	go TurnStart(game)
}
