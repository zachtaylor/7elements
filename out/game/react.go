package game

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func React(player out.Player, g *game.T, username string) { // stateid, username, react string, timer cast.Duration) {
	player.Send("/game/react", cast.JSON{
		"stateid":  g.State.ID(),
		"username": username,
		"react":    g.State.Reacts[username],
		"timer":    int(g.State.Timer.Seconds()),
	})
}
