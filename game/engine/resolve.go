package engine

import "github.com/zachtaylor/7elements/game"

func resolve(g *game.G, state *game.State) (next *game.State) {
	log := g.Log().Add("State", state)
	log.Trace("done")
	state.T.Timer = 0
	states := make([]game.Phaser, 0)
	if triggered := game.TryOnFinish(g, state); len(triggered) > 0 {
		states = append(states, triggered...)
	}
	if state.T.Stack != nil {
		log.Add("Next", state.T.Stack).Debug("stackpop")
		next = state.T.Stack
		game.TryOnConnect(g, state.T.Stack.T.Phase, nil)
	} else if nextp := state.T.Phase.Next(); nextp == nil {
		log.Debug("no next")
	} else {
		log.Add("Next", next).Debug("getnext")
		next = g.NewState(nextp.Priority()[0], game.NewStateContext(nextp))
		if triggered := game.TryOnActivate(g, nextp); len(triggered) > 0 {
			states = append(states, triggered...)
		}
	}

	next = stack(g, next, states) // stack new states

	return
}
