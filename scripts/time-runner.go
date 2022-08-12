package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
)

func init() { game.Scripts["time-runner"] = TimeRunner }

func TimeRunner(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	return trigger.DrawCard(g, g.Player(ctx.Player), 1), nil
}
