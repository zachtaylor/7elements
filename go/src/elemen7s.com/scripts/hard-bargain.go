package scripts

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"elemen7s.com/engine"
)

const HardBargainID = "hard-bargain"

func init() {
	engine.Scripts[HardBargainID] = HardBargain
}

func HardBargain(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(HardBargainID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(HardBargainID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(HardBargainID)
	} else if card.Card.CardType != vii.CTYPitem {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type item").Error(HardBargainID)
	} else {
		delete(ownerSeat.Alive, gcid)
		ownerSeat.Graveyard[gcid] = card
		// animate.BroadcastSeatUpdate(game, seat)
		animate.BroadcastSeatUpdate(game, ownerSeat)
		log.Info(HardBargainID)
	}
	return nil
}
