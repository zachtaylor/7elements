package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const IntroToNecromancyID = "intro-necromancy"

func init() { game.Scripts[IntroToNecromancyID] = IntroToNecromancy }

func IntroToNecromancy(g *game.G, ctx game.ScriptContext) (rs []game.Phaser, err error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else if target, err := target.MyPastBeing(g, player, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		tokenCtx := game.NewTokenContext(target)
		tokenCtx.Text = "A zombie!"
		tokenCtx.Body.Life = 1

		if triggered := trigger.TokenAdd(g, player, tokenCtx); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
		if triggered := trigger.PlayerDamage(g, player, 1); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}
	return
}
