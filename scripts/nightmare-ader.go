package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const nightmareaderID = "nightmare-ader"

func init() {
	game.Scripts[nightmareaderID] = NightmareAder
}

func NightmareAder(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if token, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return trigger.TokenDamage(g, token, token.T.Body.Attack), nil
	}
}
