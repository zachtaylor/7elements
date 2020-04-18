package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func resolve(g *game.T) {
	log := g.Log().Add("State", g.State).Tag("engine/resolve")
	g.State.Timer = 0
	states := finish(g)

	if g.State.Stack != nil {
		g.State = g.State.Stack
		log.Add("Next", g.State).Debug("stackpop")
	} else if stater := g.State.R.GetNext(g); stater == nil { // wat
		g.State = nil // pray that states has next state
		log.Add("State", g.State).Source().Error("wtf tho")
	} else {
		g.State = g.NewState(stater)
		states = append(states, activate(g)...)
		log.Add("Next", g.State).Debug("getnexted")
	}

	stack(g, states)
	update.State(g)
}
