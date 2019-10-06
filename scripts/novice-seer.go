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

func NoviceSeer(g *game.T, s *game.Seat, target interface{}) []game.Event {
	me, ok := target.(*game.Card)
	if !ok {
		return nil
	}
	card := s.Deck.Draw()
	return []game.Event{event.NewChoiceEvent(
		s.Username,
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
				"Seat": s.Username,
			}).Tag("/scripts/" + NoviceSeerId)

			if destroy := cast.Bool(val); destroy {
				log.Debug("destroy")
				s.Past[card.Id] = card
				g.SendAll(game.BuildSeatUpdate(s))
			} else {
				log.Debug("keep")
				s.Deck.Prepend(card)
			}
		},
	)}
}
