package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const cloningpoolID = "cloning-pool"

func init() { game.Scripts[cloningpoolID] = CloningPool }

func CloningPool(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else if target, err := target.MyPresentBeing(g, player, ctx.Targets[0]); err != nil {
		return nil, err
	} else if targetCard := g.Card(target.T.Card); targetCard == nil {
		return nil, ErrCardID
	} else {
		tokenCtx := game.NewTokenContext(targetCard)
		tokenCtx.Text = "A clone"
		return trigger.TokenAdd(g, player, tokenCtx), nil
	}
}
