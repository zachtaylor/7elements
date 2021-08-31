package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

const PantherID = "panther"

func init() {
	script.Scripts[PantherID] = Panther
}

func Panther(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if token, _ := me.(*token.T); token == nil {
		err = ErrMeToken
	} else if len(args) < 0 {
		err = ErrNoTarget
	} else if targetToken, _err := checktarget.OtherPresentBeing(game, seat, token, args[0]); _err != nil {
		err = _err
	} else {
		rs = append(rs, phase.NewCombat(seat.Username, token, targetToken))
	}
	return
}
