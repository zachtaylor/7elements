package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/request"
	"github.com/zachtaylor/7elements/game/seat"
)

func Request(game *game.T, seat *seat.T, uri string, json map[string]any) []game.Phaser {

	isSeatMain := game.Phase() == "main" && game.State.Phase.Seat() == seat.Username

	log := game.Log().With(map[string]any{
		"User":    seat,
		"URI":     uri,
		"Phase":   game.Phase(),
		"StateID": game.State.ID(),
		"data":    json, // lowercase so it logs last
	})
	log.Trace("request")

	switch uri {
	case "connect":
		request.Connect(game, seat)
	case game.State.ID():
		request.Phase(game, seat, json)
	// case "chat":
	// 	request.Chat(game, seat, json)
	case "pass":
		request.Pass(game, seat, json)
	case "trigger":
		return request.Trigger(game, seat, json)
	case "attack":
		if !isSeatMain {
			return nil
		}
		return request.Attack(game, seat, json)
	case "play":
		return request.Play(game, seat, json, !isSeatMain)
	default:
		log.Warn("not found")
	}
	return nil
}
