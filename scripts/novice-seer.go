package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const NoviceSeerId = "novice-seer"

func init() {
	game.Scripts[NoviceSeerId] = NoviceSeer
}

func NoviceSeer(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	me, ok := target.(*game.Card)
	if !ok {
		return nil
	}
	card := seat.Deck.Draw()
	return []game.Event{event.NewChoiceEvent(
		seat.Username,
		"Novice Seer",
		cast.JSON{
			"card": me.JSON(),
		},
		[]cast.JSON{
			cast.JSON{
				"choice":  "false",
				"display": "Future " + card.Card.Name,
			},
			cast.JSON{
				"choice":  "true",
				"display": card.Card.Name + " to Past",
			},
		},
		func(val interface{}) {
			log := g.Log().With(log.Fields{
				"Seat": seat.Username,
			}).Tag("/scripts/" + NoviceSeerId)

			if destroy := cast.Bool(val); destroy {
				log.Debug("destroy")
				seat.Past[card.Id] = card
				g.SendSeatUpdate(seat)
			} else {
				log.Debug("keep")
				seat.Deck.Prepend(card)
			}
		},
	)}
}
