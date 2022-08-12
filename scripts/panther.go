package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/target"
)

const PantherID = "panther"

func init() { game.Scripts[PantherID] = Panther }

func Panther(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if token := g.Token(ctx.Source); token == nil {
		return nil, ErrTokenID
	} else if target, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		return []game.Phaser{
			phase.NewCombat(game.PriorityContext(g.NewPriority(ctx.Player)), token, target),
		}, nil
	}
}
