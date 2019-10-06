package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const BurnID = "burn"

func init() {
	game.Scripts[BurnID] = Burn
}

func Burn(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": seat.Username,
	}).Tag("scripts/" + BurnID)

	gcid := cast.String(target)
	card := g.Cards[gcid]
	if card == nil {
		log.Error("gcid not found")
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Error("card owner not found")
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Error("card not in present")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Error("card not type body")
	} else {
		log.Info("confirm")
		return game.TriggerDamage(g, card, 2)
	}

	return nil
}
