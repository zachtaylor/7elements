package engine

import "github.com/zachtaylor/7elements/game"

// stack adds new States as Stacked States
func stack(g *game.G, state *game.State, rs []game.Phaser) (next *game.State) {
	log := g.Log().Add("State", state)
	log.Add("Stack", stack).Trace("root")
	next = state
	var buff []game.Phaser
	for _, r := range rs {
		next = g.NewState(r.Priority()[0], game.NewStateContext(r))
		next.T.Stack = state
		state = next

		if triggered := game.TryOnActivate(g, r); len(triggered) > 0 {
			buff = append(buff, triggered...)
		}
	}
	if len(buff) > 0 {
		log.Debug("Hold my cards, lads, I'm going in!")
		next = stack(g, next, buff)
	}
	return
}
