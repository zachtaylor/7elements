package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
)

const EnergizeID = "energize"

func init() {
	engine.Scripts[EnergizeID] = Energize
}

func Energize(game *game.T, seat *game.Seat, target interface{}) game.Event {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(EnergizeID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(EnergizeID)
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error(EnergizeID)
	} else if card.Card.Type != vii.CTYPbody && card.Card.Type != vii.CTYPitem {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type bodyoritem").Error(EnergizeID)
	} else {
		card.IsAwake = true
		if card.Username != seat.Username {
			animate.GameCard(game, card)
		}
		animate.GameSeat(game, ownerSeat)
		log.Info(EnergizeID)
	}
	return nil
}
