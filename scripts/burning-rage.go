package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const BurningRageID = "burning-rage"

func init() {
	script.Scripts[BurningRageID] = BurningRage
}

func BurningRage(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else {
		for _, name := range game.Seats.Keys() {
			if name == seat.Username {
				continue
			} else if e := trigger.DamageSeat(game, game.Seats.Get(name), 1); len(e) > 0 {
				rs = append(rs, e...)
			}
		}
	}
	return
}
