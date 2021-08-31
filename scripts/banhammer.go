package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const banhammerID = "banhammer"

func init() {
	script.Scripts[banhammerID] = Banhammer
}

func Banhammer(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if card, _err := checktarget.PastCard(game, seat, args[0]); _err != nil {
		err = _err
	} else if owner := game.Seats.Get(card.User); owner == nil {
		err = ErrBadTarget
	} else {
		delete(owner.Past, card.ID)
		game.Seats.WriteSync(wsout.GameSeatJSON(seat.Data()))
	}
	return
}
