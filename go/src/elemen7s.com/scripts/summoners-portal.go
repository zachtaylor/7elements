package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"ztaylor.me/log"
)

const SummonersPortalID = "summoners-portal"

func init() {
	games.Scripts[SummonersPortalID] = SummonersPortal
}

func SummonersPortal(game *games.Game, seat *games.Seat, target interface{}) {
	card := seat.Deck.Draw()
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Card":     card,
	})

	if card == nil {
		game.Active.Activate(game)
		log.Error(SummonersPortalID + `: card is nil`)
	} else if card.Card.CardType == vii.CTYPbody || card.Card.CardType == vii.CTYPitem {
		seat.Alive[card.Id] = card
		games.BroadcastAnimateSpawn(game, card)

		if power := card.Card.GetPlayPower(); power != nil {
			game.PowerScript(seat, power, nil)
		} else {
			game.Active.Activate(game)
		}
	} else {
		log.Add("BurnedCard", true)
		seat.Graveyard[card.Id] = card
		games.BroadcastAnimateSeatUpdate(game, seat)
		game.Active.Activate(game)
	}
	log.Info(SummonersPortalID)
}
