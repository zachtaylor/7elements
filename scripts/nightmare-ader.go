package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const nightmareaderID = "nightmare-ader"

func init() {
	script.Scripts[nightmareaderID] = NightmareAder
}

func NightmareAder(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.PresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else {
		rs = trigger.DamageToken(game, token, token.Body.Attack)
	}
	return
}
