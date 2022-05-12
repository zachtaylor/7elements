package scripts

import (
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const banhammerID = "banhammer"

func init() { game.Scripts[banhammerID] = Banhammer }

func Banhammer(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if target, err := target.PastCard(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else if owner := g.Player(target.Player()); owner == nil {
		return nil, ErrBadTarget
	} else {
		delete(owner.T.Past, target.ID())
		g.MarkUpdate(owner.ID())
		return nil, nil
	}
}
