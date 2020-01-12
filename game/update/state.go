package update

import "github.com/zachtaylor/7elements/game"

// State sends game state data using SendUpdate
func State(g *game.T) {
	Game(g, "/game/state", g.State.JSON())
}
