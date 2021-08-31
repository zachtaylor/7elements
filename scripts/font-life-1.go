package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const fontoflife1ID = "font-life-1"

func init() {
	script.Scripts[fontoflife1ID] = FontOfLife1
}

func FontOfLife1(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else {
		rs = trigger.HealSeat(game, seat, 1)
	}
	return
}
