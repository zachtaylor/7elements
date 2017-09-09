package games

import (
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func Pass(game *Game, socket *sessions.Socket, log log.Log) {
	if game.GamePhase == GPHSplay {
		if game.CurrentTurn().Username == socket.Username {
			if game.TurnPhase == TPHSplay {
				log.Info("pass: voluntary play timeout")
				game.Context.Cancel()
			} else if game.TurnPhase == TPHSattack {
				log.Info("pass: voluntary attack timeout")
				game.Context.Cancel()
			} else {
				log.Add("TurnPhase", game.TurnPhase).Warn("pass: unsupported turn phase")
			}
		} else {
		}
	} else if game.GamePhase == GPHSrespond {
		game.Context.Store("respond:"+socket.Username, "true")
		SendAllSeats(game, "timer", json.Json{
			"username": socket.Username,
			"timer":    0,
		})
		log.Info("pass: add no response to play")
		checkRespond(game)
	} else {
		log.Warn("pass: wtf")
	}
}

func checkRespond(game *Game) {
	for _, seat := range game.Seats {
		if game.Context.Get("respond:"+seat.Username) == nil {
			return
		}
	}

	game.Context.Cancel()
}
