package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const GraveBirthID = "grave-birth"

func init() { game.Scripts[GraveBirthID] = GraveBirth }

func GraveBirth(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else if target, err := target.MyPastBeing(g, player, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		tokenCtx := game.NewTokenContext(target)
		tokenCtx.Text = "Birthed from the grave"
		return trigger.TokenAdd(g, player, tokenCtx), nil
	}
}
