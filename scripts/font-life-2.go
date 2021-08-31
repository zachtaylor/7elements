package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const fontoflife2ID = "font-life-2"

func init() {
	script.Scripts[fontoflife2ID] = FontOfLife2
}

func FontOfLife2(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.MyPresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		rs = trigger.HealToken(game, token, 1)
	}
	return
}
