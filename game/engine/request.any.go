package engine

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func RequestAny(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
	g.Log().Add("Data", json).Source().Debug()
	switch uri {
	case "connect":
		requestConnect(g, seat)
	case "disconnect":
		requestDisconnect(g, seat)
	case "chat":
		requestChat(g, seat, json)
	case "pass":
		requestPass(g, seat, json)
	case g.State.ID():
		RequestGameState(g, seat, json)
	case "trigger":
		return RequestTrigger(g, seat, json)
	case "attack":
		return RequestAttack(g, seat, json)
	case "play":
		return RequestPlay(g, seat, json, false)
	default:
		g.Log().With(cast.JSON{
			"Seat":  seat.String(),
			"URI":   uri,
			"State": g.State.Print(),
		}).Warn("engine/request: 404")
	}
	return nil
}
