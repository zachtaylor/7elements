package request

import (
	"github.com/zachtaylor/7elements/game"
)

func Connect(g *game.G, state *game.State, player *game.Player) {
	// if player != nil {
	// out.SendData(player.T.Writer.Name())
	// }
	game.TryOnConnect(g, state.T.Phase, player)
}
