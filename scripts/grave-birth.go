package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

const GraveBirthID = "grave-birth"

func init() {
	script.Scripts[GraveBirthID] = GraveBirth
}

func GraveBirth(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if card, _err := checktarget.MyPastBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else if card == nil {
		err = ErrNoTarget
	} else {
		rs = trigger.NewToken(game, seat, token.New(card, seat.Username))
	}
	return
}
