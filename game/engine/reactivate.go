package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func reactivate(g *game.T) {
	g.Log().Add("State", g.State).Source().Debug()
	g.State.Timer = g.Runtime.Timeout
	g.State.Reacts = make(map[string]string)
	update.State(g)
	connect(g, nil)
}
