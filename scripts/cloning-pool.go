package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
)

const CloningPoolID = "cloning-pool"

func init() {
	engine.Scripts[CloningPoolID] = CloningPool
}

func CloningPool(g *game.T, seat *game.Seat, target interface{}) game.Event {
	log := g.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := g.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(CloningPoolID)
		return nil
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(CloningPoolID)
		return nil
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Add("Error", "card not in play").Error(CloningPoolID)
		return nil
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(CloningPoolID)
		return nil
	}

	clone := game.NewCard(card.Card)
	clone.Username = seat.Username
	g.RegisterCard(clone)
	seat.Life++
	animate.GameSeat(g, seat)

	log.Info(CloningPoolID)
	return nil
}
