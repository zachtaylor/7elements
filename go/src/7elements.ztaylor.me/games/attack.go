package games

import (
	//
	"7elements.ztaylor.me/games/cards"
	"7elements.ztaylor.me/server/sessionman"
	"ztaylor.me/ctxpert"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func Attack(game *Game, socket *sessionman.Socket, data json.Json, log log.Log) {
	if game.GamePhase != GPHSplay {
		log.Add("GamePhase", game.GamePhase).Warn("attack: gamephase rejected")
	} else if game.TurnPhase != TPHSattack {
		log.Add("TurnPhase", game.TurnPhase).Warn("attack: turnphase rejected")
	} else if seat := game.GetSeat(socket.Username); seat == nil {
		log.Error("attack: seat missing")
	} else if gcid, ok := data["gcid"].(float64); !ok {
		log.Warn("attack: gcid missing")
	} else if gcard := seat.ActiveCardGCID(int(gcid)); gcard == nil {
		log.Warn("attack: gcid invalid")
	} else if !gcard.Active {
		log.Warn("attack: card is inactive")
	} else if target, ok := data["target"].(string); !ok {
		log.Warn("attack: target missing")
	} else {
		doAttack(game, gcard, target)
		log.Add("GameClockStore", game.Context.CopyStore()).Debug("attack")
	}
}

func doAttack(game *Game, gcard *gamecards.GameCard, target string) {
	if game.Context.Get("attacks") == nil {
		game.Context.Store("attacks", make(map[int]string))
	}
	attacks := game.Context.Get("attacks").(map[int]string)
	gcard.Active = false
	attacks[gcard.GameCardId] = target
	game.Context.Store("attacks", attacks)
	game.Context = ctxpert.WithNewTimeout(game.Context, game.Patience)

	SendAllSeats(game, "attack", json.Json{
		"gcid":   gcard.GameCardId,
		"target": target,
		"timer":  game.Context.Timer().Seconds(),
	})
}
