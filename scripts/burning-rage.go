package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/log"
)

const BurningRageID = "burning-rage"

func init() {
	game.Scripts[BurningRageID] = BurningRage
}

func BurningRage(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag(logtag + BurningRageID)

	me := arg.(*game.Card)
	var events []game.Event
	for _, s := range g.Seats {
		if s == seat {
			continue
		} else if e := trigger.DamageSeat(g, me, s, 1); len(e) > 0 {
			log.Copy().Add("Seat", s.Username).Info()
			events = append(events, e...)
		}
	}
	return events
}
