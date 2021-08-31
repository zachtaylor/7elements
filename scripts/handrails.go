package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

const HandrailsID = "handrails"

func init() {
	script.Scripts[HandrailsID] = Handrails
}

func Handrails(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if token, ok := me.(*token.T); !ok || token == nil {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.MyPresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		rs = trigger.WakeToken(game, token)
	}
	return
}
