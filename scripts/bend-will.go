package scripts

import (
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const bendwillID = "bend-will"

func init() { game.Scripts[bendwillID] = BendWill }

func BendWill(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if target, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return []game.Phaser{phase.NewChoice(
			ctx.Player,
			target.T.Name+"-> Awake or Asleep?",
			map[string]any{
				"tokenid": target.ID(),
			},
			[]map[string]any{
				map[string]any{"choice": "awake", "display": `<a>Awake</a>`},
				map[string]any{"choice": "asleep", "display": `<a>Asleep</a>`},
			},
			func(val any) {
				if val == "awake" {
					target.T.Awake = true
				} else if val == "asleep" {
					target.T.Awake = false
				} else {
					return
				}
				g.MarkUpdate(target.ID())
			},
		)}, nil
	}
}
