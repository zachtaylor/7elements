package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

const cloningpoolID = "cloning-pool"

func init() {
	script.Scripts[cloningpoolID] = CloningPool
}

func CloningPool(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if target, _err := checktarget.MyPresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		target.Body.Health++
		rs = trigger.NewToken(game, seat, token.New(target.Card, seat.Username))
	}
	return
}
