package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

const LightningStrikeID = "lightning-strike"

func init() {
	engine.Scripts[LightningStrikeID] = LightningStrike
}

func LightningStrike(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(LightningStrikeID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(LightningStrikeID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(LightningStrikeID)
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(LightningStrikeID)
	} else {
		engine.Damage(game, card, 3)
		animate.GameSeat(game, seat)
		animate.GameSeat(game, ownerSeat)

		log.Info(LightningStrikeID)
	}
	return nil
}
