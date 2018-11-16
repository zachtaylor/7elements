package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

const HandrailsID = "handrails"

func init() {
	engine.Scripts[HandrailsID] = Handrails
}

func Handrails(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
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
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(HandrailsID)
		return nil
	}

	card.IsAwake = true
	animate.GameCard(game, card)
	animate.GameSeat(game, seat)
	log.Info(HandrailsID)
	return nil
}
