package games

import (
	"ztaylor.me/ctxpert"
	"ztaylor.me/events"
	"ztaylor.me/log"
)

func init() {
	events.On("GameStart", func(args ...interface{}) {
		GameStart(args[0].(*Game))
	})
}

func GameStart(game *Game) {
	if game.GamePhase != GPHSwait {
		log.Add("GameId", game.Id).Add("GamePhase", game.GamePhase).Warn("gamestart: rejected")
		return
	}

	game.GamePhase = GPHSbegin

	game.Context = ctxpert.WithNewTimeout(ctxpert.New(), game.Patience)
	game.Context.Done(func(ctx *ctxpert.Context) {
		log.Add("GameId", game.Id).Add("Responses", ctx.CopyStore()).Info("gamestart: timeout")
		HandTimeout(game)
	})
	game.Context.Always(func(ctx *ctxpert.Context) {
		game.GamePhase = GPHSplay
		log.Add("GameId", game.Id).Add("Responses", ctx.CopyStore()).Debug("gamestart: finishing...")
		TurnStart(game)
	})

	for _, seat := range game.Seats {
		DealHand(game, seat)
		seat.Life = 7
	}

	log.Add("GameId", game.Id).Debug("gamestart: started...")
}
