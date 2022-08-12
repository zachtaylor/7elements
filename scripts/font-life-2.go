package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const fontoflife2ID = "font-life-2"

func init() { game.Scripts[fontoflife2ID] = FontOfLife2 }

func FontOfLife2(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else if target, err := target.MyPresentBeing(g, player, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return trigger.TokenHeal(g, target, 1), nil
	}
}
