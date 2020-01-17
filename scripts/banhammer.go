package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/update"
)

const banhammerID = "banhammer"

func init() {
	game.Scripts[banhammerID] = Banhammer
}

func Banhammer(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if card, e := target.PastBeingItem(g, s, args[0]); e != nil {
		err = e
	} else if s := g.GetSeat(card.Username); s == nil {
		err = ErrBadTarget
	} else {
		delete(s.Past, card.ID)
		update.Seat(g, s)
	}
	return
}
