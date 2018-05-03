package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
)

const CloningPoolID = "cloning-pool"

func init() {
	games.Scripts[CloningPoolID] = CloningPool
}

func CloningPool(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(CloningPoolID)
		return
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(CloningPoolID)
		return
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(CloningPoolID)
		return
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(CloningPoolID)
		return
	}

	game.RegisterToken(seat.Username, vii.NewGameCard(card.Card, card.CardText))
	seat.Life++
	games.BroadcastAnimateSeatUpdate(game, seat)

	log.Info(CloningPoolID)
}
