package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/request"
)

func Request(game *game.G, state *game.State, player *game.Player, uri string, json map[string]any) []game.Phaser {

	log := game.Log().With(map[string]any{
		"Player":  player.ID(),
		"URI":     uri,
		"Phase":   state.T.Phase.Type(),
		"StateID": state.ID(),
		"data":    json, // lowercase so it logs last
	})
	log.Trace("request")

	switch uri {
	case "connect":
		request.Connect(game, state, player)
	// case "chat":
	// 	request.Chat(game, seat, json)
	case "pass":
		request.Pass(game, state, player, json)
	case state.ID():
		request.Phase(game, state, player, json)
	case "trigger":
		return request.Trigger(game, state, player, json)
	case "attack":
		return request.Attack(game, state, player, json)
	case "play":
		onlySpells := state.T.Phase.Type() != "main" || state.T.Phase.Priority()[0] != player.ID()
		return request.Play(game, state, player, json, onlySpells)
	default:
		log.Warn("not found")
	}
	return nil
}
