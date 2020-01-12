package update

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

// React sends JSON command to update state reaction
func React(g *game.T, username string) {
	Game(g, "/game/react", cast.JSON{
		"stateid":  g.State.ID(),
		"username": username,
		"react":    g.State.Reacts[username],
		"timer":    int(g.State.Timer.Seconds()),
	})
}
