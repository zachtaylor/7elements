package scripts

import (
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"ztaylor.me/log"
)

const BoenID = "boen"

func init() {
	engine.Scripts[BoenID] = Boen
}

func Boen(g *game.T, seat *game.Seat, target interface{}) game.Event {
	log := g.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Card":     target,
	})

	if self, ok := target.(*game.Card); !ok {
		log.Error(BoenID + `: self target failed`)
	} else if card := seat.Deck.Draw(); card == nil {
		log.Error(BoenID + `: deck is empty`)
	} else {
		self.Body.Health++
		seat.Hand[card.Id] = card

		animate.GameHand(g, seat)
		animate.GameSeat(g, seat)
	}

	log.Info(BoenID)
	return nil
}
