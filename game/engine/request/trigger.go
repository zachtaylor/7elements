package request

import (
	"reflect"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func Trigger(game *game.T, seat *seat.T, json map[string]any) (rs []game.Phaser) {
	log := game.Log().Add("Seat", seat)

	// validation

	tokenID, _ := json["id"].(string)
	if len(tokenID) < 1 {
		log.Error("id: ", json["id"])
		return
	}

	log = log.Add("TokenID", tokenID)

	token, err := target.MyPresent(game, seat, tokenID)
	if err != nil {
		log.Add("Error", err).Error()
		seat.Writer.WriteMessageData(wsout.Error("trigger", err.Error()))
		return nil
	}

	var powerID int
	if poweridbuff, _ := json["powerid"].(float64); int(poweridbuff) < 1 {
		game.Log().Add("powerid", json["powerid"]).Add("type", reflect.TypeOf(json["powerid"])).Warn("powerid missing")
		return nil
	} else {
		powerID = int(poweridbuff)
	}

	log = log.Add("PowerId", powerID)

	power := token.Powers[powerID]
	if power == nil {
		log.Add("Keys", token.Powers.Keys()).Error("power not found")
		return
	} else if !token.IsAwake && power.UsesTurn {
		log.Error("card is asleep")
		seat.Writer.WriteMessageData(wsout.Error(token.Card.Proto.Name, "not awake"))
		return
	} else if !seat.Karma.Active().Test(power.Costs) {
		log.Add("Costs", power.Costs).Error("cannot afford")
		seat.Writer.WriteMessageData(wsout.Error(token.Card.Proto.Name, "not enough elements"))
		return
	}

	targetID, _ := json["target"].(string)
	if !target.IsValid(game, seat, power.Target, targetID) {
		log.Error("targetid: ", json["target"])
	}

	seat.Karma.Deactivate(power.Costs)
	if power.Costs.Total() > 0 {
		game.Seats.WriteMessageData(wsout.GameSeat(seat.JSON()))
	}
	if power.UsesTurn {
		rs = append(rs, trigger.SleepToken(game, token)...)
	}
	if power.UsesLife {
		rs = append(rs, trigger.RemoveToken(game, token)...)
	}

	log.Add("TargetID", targetID).Add("Power", power).Trace()

	rs = append(rs, phase.NewTrigger(seat.Username, token, power, targetID))
	return
}
