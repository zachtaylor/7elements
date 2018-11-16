package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
	"ztaylor.me/log"
)

const BoenID = "boen"

func init() {
	engine.Scripts[BoenID] = Boen
}

func Boen(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Card":     target,
	})

	if self, ok := target.(*vii.GameCard); !ok {
		log.Error(BoenID + `: self target failed`)
	} else if card := seat.Deck.Draw(); card == nil {
		log.Error(BoenID + `: deck is empty`)
	} else {
		self.Body.Health++
		seat.Hand[card.Id] = card

		animate.GameHand(game, seat)
		animate.GameSeat(game, seat)
	}

	log.Info(BoenID)
	return nil
}
