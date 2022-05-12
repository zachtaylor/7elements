package scripts

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
)

func init() { game.Scripts["time-runner"] = TimeRunner }

func TimeRunner(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	return trigger.DrawCard(g, g.Player(ctx.Player), 1), nil
}
