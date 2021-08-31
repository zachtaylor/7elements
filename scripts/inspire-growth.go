package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const inspiregrowthID = "inspire-growth"

func init() {
	script.Scripts[inspiregrowthID] = InspireGrowth
}

func InspireGrowth(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := checktarget.PresentBeing(game, seat, args[0]); _err != nil {
		err = _err
	} else if token == nil {
		err = ErrNoTarget
	} else {
		token.Body.Attack++
		game.Seats.WriteSync(wsout.GameToken(token.Data()).EncodeToJSON())
	}
	return
}
