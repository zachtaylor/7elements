package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
)

const fontoflife1ID = "font-life-1"

func init() { game.Scripts[fontoflife1ID] = FontOfLife1 }

func FontOfLife1(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else {
		return trigger.PlayerHeal(g, player, 1), nil
	}
}
