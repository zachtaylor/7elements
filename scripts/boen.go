package scripts

import (
	"github.com/zachtaylor/7tcg"
	"github.com/zachtaylor/7tcg/animate"
	"github.com/zachtaylor/7tcg/engine"
	"ztaylor.me/log"
)

const BoenID = "boen"

func init() {
	engine.Scripts[BoenID] = Boen
}

func Boen(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
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

		animate.Hand(game, seat)
		animate.BroadcastSeatUpdate(game, seat)
	}

	log.Info(BoenID)
	return nil
}
