package games

// import (
// 	"github.com/zachtaylor/7tcg"
// 	"ztaylor.me/js"
// )

// func TryTrigger(game *vii.Game, seat *vii.GameSeat, json js.Object) {
// 	log := game.Log().Add("Username", seat.Username).Add("Elements", seat.Elements.String())

// 	gcid := json.Sval("gcid")
// 	if gcid == "" {
// 		log.Error("games.Trigger: gcid missing")
// 		return
// 	}

// 	powerid := json.Ival("powerid")
// 	if powerid < 1 {
// 		log.Error("games.Trigger: powerid missing")
// 		return
// 	}

// 	log.Add("GCID", gcid).Add("PowerId", powerid)
// 	card := seat.Alive[gcid]
// 	if card == nil {
// 		log.Error("games.Trigger: gcid not found")
// 	} else if power := card.Powers[powerid]; power == nil {
// 		log.Error("games.Trigger: powerid not found")
// 	} else if !card.IsAwake && power.UsesTurn {
// 		seat.SendError(card.CardText.Name, `not awake`)
// 		log.Error("games.Trigger: card is asleep")
// 	} else if !seat.Elements.GetActive().Test(power.Costs) {
// 		seat.SendError(card.CardText.Name, `not enough elements`)
// 		log.Error("games.Trigger: cannot afford")
// 	} else if power.Target == "self" {
// 		GameEngine.Trigger(game, seat, card, power, card)
// 	} else {
// 		GameEngine.Trigger(game, seat, card, power, json["target"])
// 	}
// }
