package games

import (
	"7elements.ztaylor.me/server/sessionman"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func Hand(game *Game, socket *sessionman.Socket, data json.Json, log log.Log) {
	log.Add("Choice", data["choice"])

	if game.GamePhase != GPHSbegin {
		log.Warn("hand: event rejected")
		return
	}

	seat := game.GetSeat(socket.Username)

	if seat == nil {
		log.Warn("hand: seat missing")
		return
	} else if data["choice"] == "mulligan" {
		DealHand(game, seat)
		game.Context.Store("hand:"+seat.Username, "mulligan")
	} else if data["choice"] == "keep" {
		game.Context.Store("hand:"+seat.Username, "keep")
	} else {
		log.Warn("hand: choice not recognized")
		return
	}

	socket.Send("game", MakeGameJson(game, socket.Username))

	for _, seatx := range game.Seats {
		if game.Context.Get("hand:"+seatx.Username) == nil {
			log.Add("Missing", seatx.Username).Debug("hand: accepted, still waiting response")
			return
		}
	}

	log.Info("hand: accepted, all choices made")
	game.Context.Cancel()
}
