package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
)

const WandOfSuppressionID = "wand-of-suppression"

func init() {
	engine.Scripts[WandOfSuppressionID] = WandOfSuppression
}

func WandOfSuppression(game *game.T, seat *game.Seat, target interface{}) game.Event {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(WandOfSuppressionID)
		return nil
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(WandOfSuppressionID)
		return nil
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error(WandOfSuppressionID)
		return nil
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(WandOfSuppressionID)
		return nil
	}

	card.IsAwake = false
	animate.GameCard(game, card)
	animate.GameSeat(game, seat)

	log.Info(WandOfSuppressionID)
	return nil
}
