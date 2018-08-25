package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

const IfritID = "ifrit"

func init() {
	engine.Scripts[IfritID] = Ifrit
}

func Ifrit(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	if target == "player" {
		for _, s := range game.Seats {
			if seat.Username != seat.Username {
				seat.Life--
				animate.BroadcastSeatUpdate(game, s)
				animate.BroadcastSeatUpdate(game, seat)

				log.Add("Seat", s).Info(IfritID)
				return nil
			}
		}
	}

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(IfritID)
		return nil
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(IfritID)
		return nil
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(IfritID)
		return nil
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(IfritID)
		return nil
	}

	engine.Damage(game, card, 1)
	animate.BroadcastSeatUpdate(game, game.GetSeat(card.Username))
	log.Info(IfritID)
	return nil
}
