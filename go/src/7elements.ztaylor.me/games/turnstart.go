package games

import (
	"ztaylor.me/ctxpert"
	"ztaylor.me/log"
)

func TurnStart(game *Game) {
	if game.GamePhase != GPHSplay {
		log.Add("GameId", game.Id).Add("GamePhase", game.GamePhase).Warn("turnstart: event rejected")
		return
	}
	if game.TurnPhase != TPHSwait {
		log.Add("GameId", game.Id).Add("TurnPhase", game.TurnPhase).Warn("turnstart: event rejected")
		return
	}

	game.Context = ctxpert.WithNewTimeout(ctxpert.New(), game.Patience)
	game.Context.Done(func(ctx *ctxpert.Context) {
		winnerTurn := game.BuildNextTurn()
		log.Warn("turnstart: timeout forfeit game")
		LoseGame(game, winnerTurn.Username)
		for _, seat := range game.Seats {
			if socket := seat.Socket; socket != nil {
				socket.Send("gamedone", MakeDoneJson(game, socket.Username))
			}
		}
	})
	game.Context.Always(func(ctx *ctxpert.Context) {
		go PlayStart(game)
	})

	game.TurnPhase = TPHSbegin
	turn := game.BuildNextTurn()
	seat := game.GetSeat(turn.Username)

	seat.Reactivate()
	SendAllSeats(game, "turn", MakeTurnJson(game, game.CurrentTurn()))

	log.Add("GameId", game.Id).Add("TurnId", len(game.History)).Add("Username", game.CurrentTurn().Username).Info("turnstart")
}
