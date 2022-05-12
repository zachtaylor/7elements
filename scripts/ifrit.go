package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
)

const IfritID = "ifrit"

func init() { game.Scripts[IfritID] = Ifrit }

func Ifrit(g *game.G, ctx game.ScriptContext) (rs []game.Phaser, err error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else {
		for _, playerID := range g.Players() {
			if playerID == ctx.Player {
			} else if target := g.Player(playerID); target == nil {
				return nil, ErrPlayerID
			} else if triggered := trigger.PlayerDamage(g, target, 1); len(triggered) > 0 {
				rs = append(rs, triggered...)
			}
		}
	}
	return
}
