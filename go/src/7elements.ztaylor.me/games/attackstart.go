package games

import (
	"ztaylor.me/ctxpert"
	"ztaylor.me/log"
)

func AttackStart(game *Game) {
	if game.GamePhase != GPHSplay {
		log.Add("GameId", game.Id).Add("GamePhase", game.GamePhase).Warn("attackstart: error")
		return
	} else if game.TurnPhase != TPHSplay {
		log.Add("GameId", game.Id).Add("TurnPhase", game.TurnPhase).Warn("attackstart: error")
		return
	}

	game.TurnPhase = TPHSattack

	game.Context = ctxpert.WithNewTimeout(ctxpert.New(), game.Patience)
	game.Context.Always(func(ctx *ctxpert.Context) {
		go DefendStart(game)
	})

	SendAllSeats(game, "turn", MakeTurnJson(game, game.CurrentTurn()))
	log.Add("GameId", game.Id).Info("attackstart")
}
