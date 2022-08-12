package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
)

const PainfulMemoriesID = "painful-memories"

func init() { game.Scripts[PainfulMemoriesID] = PainfulMemories }

func PainfulMemories(g *game.G, ctx game.ScriptContext) (rs []game.Phaser, err error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else {
		count := 0
		for cardID := range player.T.Past {
			if c := g.Card(cardID); c == nil {
				return nil, ErrCardID
			} else if c.T.Kind == card.Being {
				count++
			}
		}
		for _, playerID := range g.Players() {
			if playerID == ctx.Player {
			} else if target := g.Player(playerID); target == nil {
				return nil, ErrBadTarget
			} else if triggered := trigger.PlayerDamage(g, target, count); len(triggered) > 0 {
				rs = append(rs, triggered...)
			}
		}
	}
	return
}
