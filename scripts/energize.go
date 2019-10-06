package scripts

import (
	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

const EnergizeID = "energize"

func init() {
	game.Scripts[EnergizeID] = Energize
}

func Energize(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := cast.String(target)
	card := g.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(EnergizeID)
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(EnergizeID)
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error(EnergizeID)
	} else if card.Card.Type != vii.CTYPbody && card.Card.Type != vii.CTYPitem {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type bodyoritem").Error(EnergizeID)
	} else {
		card.IsAwake = true
		g.SendAll(game.BuildCardUpdate(card))
		log.Info(EnergizeID)
	}
	return nil
}
