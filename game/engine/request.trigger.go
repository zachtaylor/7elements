package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func RequestTrigger(g *game.T, seat *game.Seat, json cast.JSON) []game.Stater {
	log := g.Log().Add("Seat", seat)

	// validation

	token, err := target.MyPresent(g, seat, json.GetS("id"))
	if err != nil {
		log.Source().Add("Error", err).Source().Error()
		update.ErrorW(seat, "trigger", err.Error())
		return nil
	}

	powerid := json.GetI("powerid")
	if powerid < 1 {
		log.Source().Error("powerid missing")
		return nil
	}
	log.Add("PowerId", powerid)

	power := token.Powers[powerid]
	if power == nil {
		log.Source().Error("powerid not found")
		return nil
	} else if !token.IsAwake && power.UsesTurn {
		update.ErrorW(seat, token.Card.Card.Name, `not awake`)
		log.Source().Error("card is asleep")
		return nil
	} else if !seat.Karma.Active().Test(power.Costs) {
		update.ErrorW(seat, token.Card.Card.Name, `not enough elements`)
		log.Add("Costs", power.Costs).Source().Error("cannot afford")
		return nil
	}

	// exec
	log.Add("Token", token).Add("Power", power).Source().Debug()

	dirty := false
	if power.Costs.Count() > 0 {
		dirty = true
		seat.Karma.Deactivate(power.Costs)
	}
	if power.UsesTurn {
		token.IsAwake = false
		update.Token(g, token)
	}
	if power.UsesKill {
		dirty = true
		delete(seat.Present, token.ID)
	}
	if dirty {
		update.Seat(g, seat)
	}

	if power.Target == "self" {
		return []game.Stater{state.NewTrigger(seat.Username, token, power, token)}
	}
	return []game.Stater{state.NewTrigger(seat.Username, token, power, json["target"])}
}
