package games

import (
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func SendAllSeats(game *Game, name string, data json.Json) {
	for _, seat := range game.Seats {
		if socket := seat.Socket; socket == nil {
			log.Add("GameId", game.Id).Add("Name", name).Add("Username", seat.Username).Warn("games: send socket missing")
		} else {
			go socket.Send(name, data)
		}
	}
}
