package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

const EnergizeID = "energize"

func init() {
	engine.Scripts[EnergizeID] = Energize
}

func Energize(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(EnergizeID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(EnergizeID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(EnergizeID)
	} else if card.Card.CardType != vii.CTYPbody && card.Card.CardType != vii.CTYPitem {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type bodyoritem").Error(EnergizeID)
	} else {
		card.IsAwake = true
		if card.Username != seat.Username {
			animate.BroadcastCardUpdate(game, card)
		}
		animate.BroadcastSeatUpdate(game, ownerSeat)
		log.Info(EnergizeID)
	}
	return nil
}
