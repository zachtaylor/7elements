package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const HardBargainID = "hard-bargain"

func init() { game.Scripts[HardBargainID] = HardBargain }

func HardBargain(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if target, err := target.PresentItem(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return trigger.TokenRemove(g, target), nil
	}
}
