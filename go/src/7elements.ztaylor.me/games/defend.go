package games

import (
	"7elements.ztaylor.me/games/cards"
	"ztaylor.me/ctxpert"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func Defend(game *Game, socket *sessions.Socket, data json.Json, log log.Log) {
	if game.GamePhase != GPHSplay {
		log.Add("GamePhase", game.GamePhase).Warn("defend: rejected")
	} else if game.TurnPhase != TPHSdefend {
		log.Add("TurnPhase", game.TurnPhase).Warn("defend: rejected")
	} else if seat := game.GetSeat(socket.Username); seat == nil {
		log.Error("defend: seat missing")
	} else if gcid, ok := data["gcid"].(float64); !ok {
		log.Warn("defend: gcid missing")
	} else if gcard := seat.ActiveCardGCID(int(gcid)); gcard == nil {
		log.Add("GCID", gcid).Warn("defend: gcid invalid")
	} else if !gcard.Active {
		log.Warn("defend: card is inactive")
	} else if targetgcid, ok := data["targetgcid"].(float64); !ok {
		log.Warn("defend: targetgcid missing")
	} else if attacks := game.Context.Get("attacks").(map[int]string); attacks == nil {
		log.Warn("defend: attacks data not found")
	} else if targetPlayer := attacks[int(targetgcid)]; targetPlayer == "" {
		log.Clone().Add("Attacker", targetgcid).Warn("defend: attacker not found")
	} else if targetPlayer != seat.Username {
		log.Clone().Add("Attacker", targetgcid).Add("Target", targetPlayer).Warn("defend: not attacking you")
	} else {
		doDefend(game, gcard, int(targetgcid))
	}
}

func doDefend(game *Game, gcard *gamecards.GameCard, targetcard int) {
	if game.Context.Get("defends") == nil {
		game.Context.Store("defends", make(map[int]int))
	}
	defends := game.Context.Get("defends").(map[int]int)
	defends[gcard.GameCardId] = targetcard
	game.Context.Store("defends", defends)
	game.Context = ctxpert.WithNewTimeout(game.Context, game.Patience)

	SendAllSeats(game, "defend", json.Json{
		"gcid":       gcard.GameCardId,
		"targetgcid": targetcard,
		"timer":      game.Context.Timer().Seconds(),
	})
}
