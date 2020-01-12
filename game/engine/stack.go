package engine

import "github.com/zachtaylor/7elements/game"

func stack(g *game.T, events []game.Stater) {
	if events == nil || len(events) < 1 {
		return
	}
	g.Log().Add("State", g.State).Add("Stack", events).Source().Debug()
	next := make([]game.Stater, 0)
	for _, e := range events {
		s := g.NewState(e)
		s.Stack = g.State
		g.State = s
		if addnext := activate(g); len(addnext) > 0 {
			g.Log().Add("State", g.State).Add("Stack", events).Source().Debug("activate triggers new state")
			next = append(next, addnext...)
		}
	}
	if len(next) > 0 {
		g.Log().Add("State", g.State).Add("Stack", events).Source().Debug("hold my cards, lads, i'm going in!")
		stack(g, next)
	}
}
