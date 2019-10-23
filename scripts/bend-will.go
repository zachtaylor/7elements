package scripts

import (
	"fmt"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const BendWillID = "bend-will"

func init() {
	game.Scripts[BendWillID] = BendWill
}

func BendWill(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + BendWillID)

	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	return []game.Event{event.NewChoiceEvent(
		seat.Username,
		fmt.Sprintf("%s -> Awake or Asleep?", card.Card.Name),
		cast.JSON{
			"card": card.JSON(),
		},
		[]cast.JSON{
			cast.JSON{"choice": "awake", "display": `<a>Awake</a>`},
			cast.JSON{"choice": "asleep", "display": `<a>Asleep</a>`},
		},
		func(val interface{}) {
			if val == "awake" {
				card.IsAwake = true
			} else if val == "asleep" {
				card.IsAwake = false
			} else {
				log.Add("val", val).Error("choice?")
				return
			}
			log.Add("val", val).Info("done")
			g.SendCardUpdate(card)
		},
	)}
}
