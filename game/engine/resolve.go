package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
)

func (t *T) resolve(game *game.T) {
	log := game.Log().Add("State", game.State)
	log.Trace("done")
	game.State.Timer = 0
	states := phase.TryOnFinish(game) // combine states
	if game.State.Stack != nil {
		log.Add("Next", game.State.Stack).Debug("stackpop")
		game.State = game.State.Stack
		phase.TryOnConnect(game, nil)
	} else if next := game.State.Phase.GetNext(game); next == nil {
		game.State = nil
		return // that's a wrap
	} else {
		log.Add("Next", next).Debug("getnext")
		game.State = t.NewState(game, next)
		states = append(states, phase.TryOnActivate(game)...) // combine states
	}

	t.stack(game, states) // stack new states
}
