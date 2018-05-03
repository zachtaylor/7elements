package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"ztaylor.me/log"
)

const BoenID = "boen"

func init() {
	games.Scripts[BoenID] = Boen
}

func Boen(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Card":     target,
	})

	if self, ok := target.(*vii.GameCard); !ok {
		log.Error(BoenID + `: self target failed`)
	} else if card := seat.Deck.Draw(); card == nil {
		log.Error(BoenID + `: deck is empty`)
	} else {
		self.CardBody.Health++
		seat.Hand[card.Id] = card

		games.AnimateHand(game, seat)
		games.BroadcastAnimateSeatUpdate(game, seat)
	}

	game.Active.Activate(game)
	log.Info(BoenID)
}
