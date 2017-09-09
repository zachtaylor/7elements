package games

import (
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func WebsocketMessage(socket *sessions.Socket, name string, data json.Json) {
	log := log.Add("Username", socket.Username).Add("Data", data)

	if name == "join" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wsjoin: missing gameid")
			return
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wsjoin: gameid is invalid")
			return
		} else {
			if seat := game.GetSeat(socket.Username); seat == nil {
				log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsjoin: not participating in game")
			} else if seat.Socket != nil {
				log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsjoin: seat already connected")
			} else {
				seat.Socket = socket
				log.Add("Username", socket.Username).Add("GameId", game.Id).Info("wsjoin: seat connected")
			}
			WebsocketJoinGame(socket, game)
		}
	} else if name == "hand" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wshand: missing gameid")
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wshand: gameid is invalid")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wshand: not participating in game")
		} else {
			Hand(game, socket, data, log)
		}
	} else if name == "element" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wselement: missing gameid")
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wselement: gameid is invalid")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wselement: not participating in game")
		} else {
			Element(game, socket, data, log)
		}
	} else if name == "play" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wsplay: missing gameid")
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wsplay: gameid is invalid")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsplay: not participating in game")
		} else {
			Play(game, socket, data, log)
		}
	} else if name == "pass" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wspass: missing gameid")
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", gameid).Warn("wspass: gameid is invalid")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wspass: not participating in game")
		} else {
			Pass(game, socket, log)
		}
	} else if name == "attack" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wsattack: missing gameid")
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsattack: gameid is invalid")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsattack: not participating in game")
		} else {
			Attack(game, socket, data, log)
		}
	} else if name == "defend" {
		if gameid, ok := data["gameid"].(float64); !ok {
			log.Warn("wsdefend: missing gameid")
		} else if game := Cache[int(gameid)]; game == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsdefend: gameid is invalid")
		} else if seat := game.GetSeat(socket.Username); seat == nil {
			log.Add("Username", socket.Username).Add("GameId", game.Id).Warn("wsdefend: not participating in game")
		} else {
			Defend(game, socket, data, log)
		}
	} else {
		log.Add("Name", name).Warn("websocket event name not recognized")
	}
}
