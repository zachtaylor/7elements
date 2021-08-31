package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

const PixieID = "pixie"

func init() {
	script.Scripts[PixieID] = Pixie
}

func Pixie(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if token, ok := me.(*token.T); !ok || token == nil {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if targetSeat := game.Seats.Get(args[0]); targetSeat != nil {
		rs = trigger.HealSeat(game, targetSeat, token.Body.Health)
		rs = append(rs, trigger.RemoveToken(game, token)...)
	} else if targetToken, _err := checktarget.OtherPresentBeing(game, seat, token, args[0]); _err != nil {
		err = _err
	} else {
		rs = trigger.HealToken(game, targetToken, token.Body.Health)
		rs = append(rs, trigger.RemoveToken(game, token)...)
	}
	return
}
