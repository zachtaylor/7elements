package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const GraceID = "grace"

func init() {
	script.Scripts[GraceID] = Grace
}

func Grace(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.MyPresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		rs = trigger.HealToken(game, token, 3)
	}
	return
}
