package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
)

func stackbuff() []game.Phaser { return nil }

// stack adds new States as Stacked States
func (t *T) stack(game *game.T, stack []game.Phaser) {
	log := game.Log().Add("State", game.State)
	log.Add("Stack", stack).Trace("root")
	next := stackbuff()
	for _, r := range stack {
		state := t.NewState(game, r)
		state.Stack = game.State
		game.State = state

		if addnext := phase.TryOnActivate(game); len(addnext) > 0 {
			game.Log().With(map[string]any{
				"State": game.State,
				"Stack": addnext,
			}).Debug("activate trigger")
			next = append(next, addnext...)
		}
	}
	if len(next) > 0 {
		log.Debug("Hold my cards, lads, I'm going in!")
		t.stack(game, next)
	}
}
