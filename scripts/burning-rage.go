package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
)

const BurningRageID = "burning-rage"

func init() {
	game.Scripts[BurningRageID] = BurningRage
}

func BurningRage(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	token, ok := me.(*game.Token)
	if !ok || token == nil {
		err = ErrMeToken
	} else {
		for _, seat := range g.Seats {
			if s == seat {
				continue
			} else if e := trigger.DamageSeat(g, token.Card, seat, 1); len(e) > 0 {
				events = append(events, e...)
			}
		}
	}
	return
}
