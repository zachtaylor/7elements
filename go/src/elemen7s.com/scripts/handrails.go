package scripts

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"elemen7s.com/engine"
)

const HandrailsID = "handrails"

func init() {
	engine.Scripts[HandrailsID] = Handrails
}

func Handrails(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(HandrailsID)
		return nil
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(HandrailsID)
		return nil
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(HandrailsID)
		return nil
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(HandrailsID)
		return nil
	}

	card.IsAwake = true
	animate.BroadcastCardUpdate(game, card)
	animate.BroadcastSeatUpdate(game, seat)
	log.Info(HandrailsID)
	return nil
}
