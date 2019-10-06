package scripts

import (
	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const GraveBirthID = "grave-birth"

func init() {
	game.Scripts[GraveBirthID] = GraveBirth
}

func GraveBirth(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Seat": seat.Username,
	}).Tag(GraveBirthID)
	if gcid := cast.String(target); gcid == "" {
		log.Warn("choice not found")
	} else if card := g.Cards[gcid]; card == nil {
		log.Warn("gcid not found")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Warn("not type body")
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("CardOwner", card.Username).Warn("card owner not found")
	} else if !ownerSeat.HasPastCard(gcid) {
		log.Add("CardOwner", card.Username).Add("Past", ownerSeat.Past.String()).Warn("card not in past")
	} else {
		log.Add("CardId", card.Card.Id).Info("confirm")
	}
	return nil
}
