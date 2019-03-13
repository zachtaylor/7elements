package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
)

const HardBargainID = "hard-bargain"

func init() {
	engine.Scripts[HardBargainID] = HardBargain
}

func HardBargain(game *game.T, seat *game.Seat, target interface{}) game.Event {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(HardBargainID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(HardBargainID)
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error(HardBargainID)
	} else if card.Card.Type != vii.CTYPitem {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type item").Error(HardBargainID)
	} else {
		delete(ownerSeat.Present, gcid)
		ownerSeat.Past[gcid] = card
		// animate.GameSeat(game, seat)
		animate.GameSeat(game, ownerSeat)
		log.Info(HardBargainID)
	}
	return nil
}
