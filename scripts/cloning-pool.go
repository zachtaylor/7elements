package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

const CloningPoolID = "cloning-pool"

func init() {
	engine.Scripts[CloningPoolID] = CloningPool
}

func CloningPool(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(CloningPoolID)
		return nil
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(CloningPoolID)
		return nil
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(CloningPoolID)
		return nil
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(CloningPoolID)
		return nil
	}

	clone := vii.NewGameCard(card.Card, card.CardText)
	clone.Username = seat.Username
	game.RegisterCard(clone)
	seat.Life++
	animate.BroadcastSeatUpdate(game, seat)

	log.Info(CloningPoolID)
	return nil
}
