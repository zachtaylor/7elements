package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

func init() {
	script.Scripts["time-runner"] = TimeRunner
}

func TimeRunner(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	rs = trigger.DrawCard(game, seat, 1)
	return
}
