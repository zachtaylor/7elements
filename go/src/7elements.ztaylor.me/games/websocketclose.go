package games

import (
	"7elements.ztaylor.me/server/sessionman"
	"ztaylor.me/log"
)

func WebsocketClose(socket *sessionman.Socket) {
	for gameid, ok := range GetActiveGames(socket.Username) {
		if !ok {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wsclose: game id error")
		} else if game := Cache[gameid]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wsclose: game missing")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wsclose: seat missing")
		} else if seat.Socket == socket {
			seat.Socket = nil
			log.Add("Username", socket.Username).Add("GameId", gameid).Info("wsclose: remove socket from gameseat")
		}
	}
}
