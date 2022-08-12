package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const WaterDancerID = "water-dancer"

func init() { game.Scripts[WaterDancerID] = WaterDancer }

func WaterDancer(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if token, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return trigger.TokenSleep(g, token), nil
	}
}
