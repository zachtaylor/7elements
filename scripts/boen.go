package scripts

// import (
// 	"github.com/zachtaylor/7elements/buildjson"
// 	"github.com/zachtaylor/7elements/game/v2"
// 	"github.com/zachtaylor/7elements/game/engine"
// 	"ztaylor.me/log"
// )

// const BoenID = "boen"

// func init() {
// 	game.Scripts[BoenID] = Boen
// }

// func Boen(g *game.T, seat *seat.T, target interface{}) []game.Phaser {
// 	log := g.Log().With(log.Fields{
// 		"Username": seat.Username,
// 		"Card":     target,
// 	})

// 	if self, ok := target.(*game.Card); !ok {
// 		log.Error(BoenID + `: self target failed`)
// 	} else if card := seat.Deck.Draw(); card == nil {
// 		log.Error(BoenID + `: deck is empty`)
// 	} else {
// 		self.Body.Health++
// 		seat.Hand[card.Id] = card

// 		seat.Send(game.BuildHandUpdate(seat))
// 		g.SendSeatUpdate(seat)
// 	}

// 	log.Info(BoenID)
// 	return nil
// }
