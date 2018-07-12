package scripts

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"elemen7s.com/engine"
)

const LightningStrikeID = "lightning-strike"

func init() {
	engine.Scripts[LightningStrikeID] = LightningStrike
}

func LightningStrike(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(LightningStrikeID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(LightningStrikeID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(LightningStrikeID)
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(LightningStrikeID)
	} else {
		engine.Damage(game, card, 3)
		animate.BroadcastSeatUpdate(game, seat)
		animate.BroadcastSeatUpdate(game, ownerSeat)

		log.Info(LightningStrikeID)
	}
	return nil
}
