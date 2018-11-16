package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

const HardBargainID = "hard-bargain"

func init() {
	engine.Scripts[HardBargainID] = HardBargain
}

func HardBargain(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(HardBargainID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(HardBargainID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(HardBargainID)
	} else if card.Card.Type != vii.CTYPitem {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type item").Error(HardBargainID)
	} else {
		delete(ownerSeat.Alive, gcid)
		ownerSeat.Graveyard[gcid] = card
		// animate.GameSeat(game, seat)
		animate.GameSeat(game, ownerSeat)
		log.Info(HardBargainID)
	}
	return nil
}
