package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func trigger(g *game.T, seat *game.Seat, json cast.JSON) []game.Stater {
	log := g.Log().Add("Seat", seat)

	// validation

	token, err := target.MyPresent(g, seat, json.GetS("id"))
	if err != nil {
		log.Add("Error", err).Error()
		out.Error(seat.Player, "trigger", err.Error())
		return nil
	}

	powerid := json.GetI("powerid")
	if powerid < 1 {
		log.Error("powerid missing")
		return nil
	}
	log.Add("PowerId", powerid)

	power := token.Powers[powerid]
	if power == nil {
		log.Error("powerid not found")
		return nil
	} else if !token.IsAwake && power.UsesTurn {
		log.Error("card is asleep")
		out.Error(seat.Player, token.Card.Proto.Name, `not awake`)
		return nil
	} else if !seat.Karma.Active().Test(power.Costs) {
		log.Add("Costs", power.Costs).Error("cannot afford")
		out.Error(seat.Player, token.Card.Proto.Name, `not enough elements`)
		return nil
	}

	// approved
	log.Add("Token", token).Add("Power", power).Trace()
	return g.Settings.Engine.TriggerTokenPower(g, seat, token, power, json["target"])
}
