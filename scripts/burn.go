package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const burnID = "burn"

func init() {
	script.Scripts[burnID] = Burn
}

func Burn(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if targetSeat := game.Seats.Get(args[0]); targetSeat != nil {
		rs = trigger.DamageSeat(game, targetSeat, 2)
	} else if token, _err := checktarget.PresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		rs = trigger.DamageToken(game, token, 2)
	}
	return
}
