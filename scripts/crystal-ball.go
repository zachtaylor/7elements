package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const CrystalBallID = "crystal-ball"

func init() {
	game.Scripts[CrystalBallID] = CrystalBall
}

func CrystalBall(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag(logtag + CrystalBallID)

	card := seat.Deck.Cards[0]

	return []game.Event{event.NewChoiceEvent(
		seat.Username,
		"Shuffle your Future?",
		cast.JSON{
			"card": card.JSON(),
		},
		[]cast.JSON{
			cast.JSON{"choice": true, "display": `<a>Yes</a>`},
			cast.JSON{"choice": false, "display": `<a>No</a>`},
		},
		func(val interface{}) {
			if cast.Bool(val) {
				log.Add("Shuffle", "Yes").Info()
				seat.Deck.Shuffle()
			} else {
				log.Add("Shuffle", "No").Info()
			}
		},
	)}
}
