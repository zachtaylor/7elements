package engine

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func RequestResponse(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
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
	case "play":
		return RequestPlay(g, seat, json, true)
	default:
		g.Log().Add("Data", json).Source().Warn("404")
	}
	return nil
}
