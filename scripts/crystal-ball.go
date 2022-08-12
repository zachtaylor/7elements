package scripts

import (
	"reflect"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
)

const crystalballID = "crystal-ball"

func init() { game.Scripts[crystalballID] = CrystalBall }

func CrystalBall(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else if len(player.T.Future) < 1 {
		return nil, ErrFutureEmpty
	} else if card := g.Card(player.T.Future[0]); card == nil {
		return nil, ErrBadTarget
	} else {
		return []game.Phaser{phase.NewChoice(
			ctx.Player,
			"Shuffle your Future?",
			map[string]any{"card": card.T.Data()},
			[]map[string]any{
				{"choice": true, "display": `<a>Yes</a>`},
				{"choice": false, "display": `<a>No</a>`},
			},
			func(val any) {
				g.Log().Add("val", val).Add("type", reflect.TypeOf(val)).Info()
				if choice, ok := val.(bool); ok {
					if choice {
						player.T.Future = deck.Shuffle(player.T.Future)
					}
				}
			},
		)}, nil
	}
}
