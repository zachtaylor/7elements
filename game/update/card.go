package update

import (
	"github.com/zachtaylor/7elements/game"
)

func Card(g *game.T, c *game.Card) {
	Game(g, "/game/card", c.JSON())
}
