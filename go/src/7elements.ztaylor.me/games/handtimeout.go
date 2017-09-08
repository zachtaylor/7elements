package games

import (
	"ztaylor.me/log"
)

func HandTimeout(game *Game) {
	log := log.Add("GameId", game.Id)

	for _, seat := range game.Seats {
		log = log.Clone().Add("Username", seat.Username)

		if v := game.Context.Get("hand:" + seat.Username); v != nil {
			log.Add("Hand", v).Debug("handtimeout: " + seat.Username + " already completed")
		} else if socket := seat.Socket; socket == nil {
			log.Error("handtimeout: session not found")
			go ForfeitGame(game, seat.Username)
		} else {
			socket.Send("game", MakeGameJson(game, seat.Username))
			log.Info("handtimeout: missing response")
		}
	}

	log.Add("GameId", game.Id).Info("handtimeout")
}
