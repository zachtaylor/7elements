package engine

import "github.com/zachtaylor/7elements/game"

func finish(g *game.T) []game.Stater {
	log := g.Log().Add("State", g.State).Source()
	if finisher, _ := g.State.R.(game.FinishStater); finisher != nil {
		log.Debug()
		return finisher.Finish(g)
	}
	log.Debug("empty")
	return nil
}
