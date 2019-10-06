package scripts

import (
	"fmt"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game/event"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const PixieID = "pixie"

func init() {
	game.Scripts[PixieID] = Pixie
}

func Pixie(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag("scripts/" + PixieID)
	pixie, ok := target.(*game.Card)
	if !ok {
		log.Add("Pixie", target).Error("target")
		return nil
	}
	hp := pixie.Body.Health
	pixie.Body.Health = 0
	log.Add("Heal", hp)
	return []game.Event{event.NewTargetEvent(
		seat.Username,
		"player-being",
		fmt.Sprintf("Target Player or Being gains %d Health", hp),
		func(target string) []game.Event {
			if target == seat.Username {
				seat.Life += hp
				g.SendAll(game.BuildSeatUpdate(seat))
				return nil
			} else if opponent := g.GetOpponentSeat(seat.Username); target == opponent.Username {
				opponent.Life += hp
				g.SendAll(game.BuildSeatUpdate(opponent))
				return nil
			}

			gcid := cast.String(target)
			card := g.Cards[gcid]
			if card == nil {
				log.Error("gcid not found")
			} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
				log.Error("card owner not found")
			} else if !ownerSeat.HasPresentCard(gcid) {
				log.Error("card not in present")
			} else if card.Card.Type != vii.CTYPbody {
				log.Add("CardType", card.Card.Type).Error("card not type body")
			} else {
				card.Body.Health += hp
				g.SendAll(game.BuildCardUpdate(card))
				log.Add("Target", card).Info()
			}
			return nil
		},
	)}
}
