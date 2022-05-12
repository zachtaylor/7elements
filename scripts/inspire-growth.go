package scripts

import (
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const inspiregrowthID = "inspire-growth"

func init() { game.Scripts[inspiregrowthID] = InspireGrowth }

func InspireGrowth(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if target, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		target.T.Body.Attack++
		g.MarkUpdate(target.ID())
		return nil, nil
	}
}
