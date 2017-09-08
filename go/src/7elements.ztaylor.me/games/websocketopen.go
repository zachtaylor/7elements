package games

import (
	"7elements.ztaylor.me/server/sessionman"
	// "ztaylor.me/json"
	"ztaylor.me/log"
)

func WebsocketOpen(socket *sessionman.Socket) {
	log.Add("Username", socket.Username).Info("wsopen")
}

func WebsocketJoinGame(socket *sessionman.Socket, game *Game) {
	log := log.Add("Username", socket.Username).Add("GameId", game.Id).Add("GamePhase", game.GamePhase)

	if game.GamePhase == GPHSdone {
		socket.Send("gamedone", MakeDoneJson(game, socket.Username))
		log.Warn("websocketopen: game is finished TODO")
	} else if game.GamePhase == GPHSbegin {
		socket.Send("gamestart", MakeGameJson(game, socket.Username))
		log.Info("websocketopen: gamestart")
	} else if game.GamePhase == GPHSplay {
		socket.Send("game", MakeGameJson(game, socket.Username))
		socket.Send("turn", MakeTurnJson(game, game.CurrentTurn()))
	} else if game.GamePhase == GPHSrespond {
		socket.Send("game", MakeGameJson(game, socket.Username))
	}

	log.Info("game joined")
}
