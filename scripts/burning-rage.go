package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
)

const BurningRageID = "burning-rage"

func init() { game.Scripts[BurningRageID] = BurningRage }

func BurningRage(g *game.G, ctx game.ScriptContext) (rs []game.Phaser, err error) {
	for _, playerID := range g.Players() {
		if playerID != ctx.Player {
			if triggered := trigger.PlayerDamage(g, g.Player(playerID), 1); len(triggered) > 0 {
				rs = append(rs, triggered...)
			}
		}
	}
	return
}
