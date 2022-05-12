package scripts

import (
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/v2"
)

const CounterID = "counter"

func init() { game.Scripts[CounterID] = Counter }

func Counter(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if state := g.State(ctx.Targets[0]); state == nil {
		return nil, ErrBadTarget
	} else if play, _ := state.T.Phase.(*phase.Play); play == nil {
		return nil, ErrBadTarget
	} else {
		play.IsCancelled = true
		return nil, nil
	}
}
