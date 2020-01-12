package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

const noviceseerId = "novice-seer"

func init() {
	game.Scripts[noviceseerId] = NoviceSeer
}

func NoviceSeer(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	card := s.Deck.Draw()
	events = []game.Stater{state.NewChoice(
		s.Username,
		"Novice Seer",
		cast.JSON{
			// "card": me.JSON(),
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
			if destroy := cast.Bool(val); destroy {
				s.Past[card.ID] = card
			} else {
				s.Deck.Prepend(card)
			}
			update.Seat(g, s)
		},
	)}
	return
}
