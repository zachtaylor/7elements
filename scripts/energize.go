package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const energizeID = "energize"

func init() { game.Scripts[energizeID] = Energize }

func Energize(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if target, err := target.PresentBeingItem(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return trigger.TokenAwake(g, target), nil
	}
}
