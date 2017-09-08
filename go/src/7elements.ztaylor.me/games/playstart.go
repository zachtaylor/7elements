package games

import (
	"ztaylor.me/ctxpert"
	"ztaylor.me/log"
)

func PlayStart(game *Game) {
	log := log.Add("GameId", game.Id)

	if game.GamePhase != GPHSplay {
		log.Add("GamePhase", game.GamePhase).Warn("playstart: event rejected")
		return
	}
	if game.TurnPhase != TPHSbegin {
		log.Add("TurnPhase", game.TurnPhase).Warn("playstart: event rejected")
		return
	}

	game.TurnPhase = TPHSplay

	game.Context = ctxpert.WithNewTimeout(ctxpert.New(), game.Patience)
	game.Context.Always(func(ctx *ctxpert.Context) {
		go AttackStart(game)
	})

	SendAllSeats(game, "turn", MakeTurnJson(game, game.CurrentTurn()))
	log.Info("playstart")
}
