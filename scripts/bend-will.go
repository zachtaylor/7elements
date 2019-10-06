package scripts

import (
	"fmt"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const BendWillID = "bend-will"

func init() {
	game.Scripts[BendWillID] = BendWill
}

func BendWill(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	gcid := cast.String(target)
	card := g.Cards[gcid]
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": seat.Username,
	}).Tag("scripts/bend-will")
	if card == nil {
		log.Add("Error", "gcid not found").Error()
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error()
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error()
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error()
	} else {
		log.Add("Card", card.Print()).Info("choice")
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
				g.SendAll(game.BuildCardUpdate(card))
			},
		)}
	}
	return nil
}
