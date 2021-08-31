package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const IfritID = "ifrit"

func init() {
	script.Scripts[IfritID] = Ifrit
}

func Ifrit(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else {
		enemy := game.Seats.GetOpponent(seat.Username)
		rs = trigger.DamageSeat(game, enemy, 1)
	}
	return
}
