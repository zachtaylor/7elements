package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const PixieID = "pixie"

func init() { game.Scripts[PixieID] = Pixie }

func Pixie(g *game.G, ctx game.ScriptContext) (rs []game.Phaser, err error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if token := g.Token(ctx.Source); token == nil {
		return nil, ErrTokenID
	} else if targetPlayer := g.Player(ctx.Targets[0]); targetPlayer != nil {
		if triggered := trigger.PlayerHeal(g, targetPlayer, token.T.Body.Life); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
		if triggered := trigger.TokenRemove(g, token); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	} else if targetToken, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		if triggered := trigger.TokenHeal(g, targetToken, token.T.Body.Life); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
		if triggered := trigger.TokenRemove(g, token); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}
	return
}
