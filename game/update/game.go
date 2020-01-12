package update

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

// Game sends data into the Game
func Game(g *game.T, uri string, data cast.JSON) {
	g.WriteJSON(Build(uri, data))
}
