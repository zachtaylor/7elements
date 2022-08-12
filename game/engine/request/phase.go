package request

import "github.com/zachtaylor/7elements/game"

func Phase(g *game.G, state *game.State, player *game.Player, json map[string]any) {
	game.TryOnRequest(g, state, player, json)
}
