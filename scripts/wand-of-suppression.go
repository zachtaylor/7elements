package scripts

// import (
// 	vii "github.com/zachtaylor/7elements"

// 	"github.com/zachtaylor/7elements/game"
// 	"ztaylor.me/cast"
// )

// const WandOfSuppressionID = "wand-of-suppression"

// func init() {
// 	game.Scripts[WandOfSuppressionID] = WandOfSuppression
// }

// func WandOfSuppression(g *game.T, seat *seat.T, target interface{}) []game.Phaser {
// 	log := g.Log().Add("Target", target).Add("Username", seat.Username)

// 	gcid := cast.String(target)
// 	card := g.Cards[gcid]
// 	if card == nil {
// 		log.Add("Error", "gcid not found").Error(WandOfSuppressionID)
// 		return nil
// 	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
// 		log.Add("Error", "card owner not found").Error(WandOfSuppressionID)
// 		return nil
// 	} else if !ownerSeat.HasPresentCard(gcid) {
// 		log.Add("Error", "card not in play").Error(WandOfSuppressionID)
// 		return nil
// 	} else if card.Card.Type != vii.CTYPbody {
// 		log.Add("CardType", card.Card.Type).Add("Error", "card not type body").Error(WandOfSuppressionID)
// 		return nil
// 	}

// 	card.IsAwake = false
// 	g.SendAll(game.BuildCardUpdate(card))
// 	g.SendSeatUpdate(seat)

// 	log.Info(WandOfSuppressionID)
// 	return nil
// }
