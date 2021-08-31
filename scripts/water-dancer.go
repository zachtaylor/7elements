package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const WaterDancerID = "water-dancer"

func init() {
	script.Scripts[WaterDancerID] = WaterDancer
}

func WaterDancer(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.PresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else if token == nil {
		err = ErrBadTarget
	} else {
		rs = trigger.SleepToken(game, token)
	}
	return
}
