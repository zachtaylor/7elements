package update

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

func Card(g *game.T, c *card.T) {
	Game(g, "/game/card", c.JSON())
}
