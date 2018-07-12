package scripts

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"elemen7s.com/engine"
	"ztaylor.me/log"
)

const SummonersPortalID = "summoners-portal"

func init() {
	engine.Scripts[SummonersPortalID] = SummonersPortal
}

func SummonersPortal(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	card := seat.Deck.Draw()
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Card":     card,
	})

	if card == nil {
		log.Error(SummonersPortalID + `: card is nil`)
	} else if card.Card.CardType == vii.CTYPbody || card.Card.CardType == vii.CTYPitem {
		seat.Alive[card.Id] = card
		animate.Spawn(game, card)

		if power := card.Card.GetPlayPower(); power != nil {
			engine.Script(game, t, seat, power, target)
		}
	} else {
		log.Add("BurnedCard", true)
		seat.Graveyard[card.Id] = card
		animate.BroadcastSeatUpdate(game, seat)
	}
	log.Info(SummonersPortalID)
	return nil
}
