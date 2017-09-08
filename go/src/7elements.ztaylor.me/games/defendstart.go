package games

import (
	"time"
	"ztaylor.me/ctxpert"
	"ztaylor.me/log"
)

func DefendStart(game *Game) {
	if game.GamePhase != GPHSplay {
		log.Add("GameId", game.Id).Warn("attacktimeout: error")
		return
	} else if game.TurnPhase != TPHSattack {
		log.Add("GameId", game.Id).Add("TurnPhase", game.TurnPhase).Warn("attacktimeout: error")
		return
	}

	game.TurnPhase = TPHSdefend

	if attacks, ok := game.Context.Get("attacks").(map[int]string); !ok {
		log.Add("GameId", game.Id).Add("Attacks", attacks).Info("defendstart: attacks missing")
		game.Context = ctxpert.WithNewTimeout(ctxpert.New(), 4*time.Second)
		game.Context.Always(func(ctx *ctxpert.Context) {
			DefendTimeout(game)
		})
	} else {
		log.Add("GameId", game.Id).Add("Attacks", attacks).Warn("defendstart: found attacks")
		game.Context = ctxpert.WithNewTimeout(ctxpert.New(), game.Patience)
		game.Context.Always(func(ctx *ctxpert.Context) {
			DefendTimeout(game)
		})
		game.Context.Store("attacks", attacks)
	}

	SendAllSeats(game, "turn", MakeTurnJson(game, game.CurrentTurn()))
	log.Add("GameId", game.Id).Info("defendstart")
}
