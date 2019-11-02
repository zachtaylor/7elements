package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const SymbiosisID = "symbiosis"

func init() {
	game.Scripts[SymbiosisID] = Symbiosis
}

func Symbiosis(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag(logtag + SymbiosisID)

	return []game.Event{event.NewTargetEvent(
		seat.Username,
		"targer-being",
		"Target Being gains 1 Attack",
		func(val string) []game.Event {
			card, err := target.PresentBeing(g, seat, val)
			if err != nil {

			} else {
				log.Add("Card", card.String()).Info()
				card.Body.Attack++
				g.SendCardUpdate(card)
			}
			return nil
		},
	)}
}
