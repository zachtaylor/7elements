package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
)

const WandOfSuppressionID = "wand-of-suppression"

func init() {
	games.Scripts[WandOfSuppressionID] = WandOfSuppression
}

func WandOfSuppression(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(WandOfSuppressionID)
		return
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(WandOfSuppressionID)
		return
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(WandOfSuppressionID)
		return
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(WandOfSuppressionID)
		return
	}

	card.IsAwake = false
	games.BroadcastAnimateCardUpdate(game, card)
	games.BroadcastAnimateSeatUpdate(game, seat)
	game.Active.Activate(game)
	log.Info(WandOfSuppressionID)
}
