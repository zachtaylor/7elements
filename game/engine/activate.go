package engine

import "github.com/zachtaylor/7elements/game"

func activate(g *game.T) (events []game.Stater) {
	if activator, ok := g.State.R.(game.ActivateStater); ok {
		g.Log().Add("State", g.State).Source().Debug()
		events = activator.OnActivate(g)
	} else {
		g.Log().Add("State", g.State).Source().Debug("none")
	}
	reactivate(g)
	return
}
