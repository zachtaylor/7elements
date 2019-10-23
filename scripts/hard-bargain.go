package scripts

// import (
// 	vii "github.com/zachtaylor/7elements"

// 	"github.com/zachtaylor/7elements/game"
// 	"ztaylor.me/cast"
// )

// const HardBargainID = "hard-bargain"

// func init() {
// 	game.Scripts[HardBargainID] = HardBargain
// }

// func HardBargain(g *game.T, seat *game.Seat, target interface{}) []game.Event {
// 	log := g.Log().Add("Target", target).Add("Username", seat.Username)

// 	gcid := cast.String(target)
// 	card := g.Cards[gcid]
// 	if card == nil {
// 		log.Add("Error", "gcid not found").Error(HardBargainID)
// 	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
// 		log.Add("Error", "card owner not found").Error(HardBargainID)
// 	} else if !ownerSeat.HasPresentCard(gcid) {
// 		log.Add("Error", "card not in play").Error(HardBargainID)
// 	} else if card.Card.Type != vii.CTYPitem {
// 		log.Add("CardType", card.Card.Type).Add("Error", "card not type item").Error(HardBargainID)
// 	} else {
// 		delete(ownerSeat.Present, gcid)
// 		g.SendAll(game.BuildCardUpdate(card))
// 		log.Info(HardBargainID)
// 	}
// 	return nil
// }
