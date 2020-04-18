package engine

import "github.com/zachtaylor/7elements/game"

func finish(g *game.T) []game.Stater {
	log := g.Log().Add("State", g.State)
	if finisher, _ := g.State.R.(game.FinishStater); finisher != nil {
		log.Source().Debug()
		return finisher.Finish(g)
	}
	log.Source().Debug("empty")
	return nil
}
