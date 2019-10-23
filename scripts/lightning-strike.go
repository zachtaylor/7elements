package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
)

const LightningStrikeID = "lightning-strike"

func init() {
	game.Scripts[LightningStrikeID] = LightningStrike
}

func LightningStrike(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := cast.String(target)
	card := g.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(LightningStrikeID)
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(LightningStrikeID)
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error(LightningStrikeID)
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(LightningStrikeID)
	} else {
		log.Info(LightningStrikeID)
		return trigger.Damage(g, card, 3)
	}
	return nil
}
