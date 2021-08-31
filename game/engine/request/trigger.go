package request

import (
	"reflect"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Trigger(game *game.T, seat *seat.T, json websocket.MsgData) (rs []game.Phaser) {
	log := game.Log().Add("Seat", seat)

	// validation

	tokenID, _ := json["id"].(string)
	if len(tokenID) < 1 {
		log.Error("id: ", json["id"])
		return
	}

	log = log.Add("TokenID", tokenID)

	token, err := checktarget.MyPresent(game, seat, tokenID)
	if err != nil {
		log.Add("Error", err).Error()
		seat.Writer.Write(wsout.ErrorJSON("trigger", err.Error()))
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
		seat.Writer.Write(wsout.ErrorJSON(token.Card.Proto.Name, "not awake"))
		return
	} else if !seat.Karma.Active().Test(power.Costs) {
		log.Add("Costs", power.Costs).Error("cannot afford")
		seat.Writer.Write(wsout.ErrorJSON(token.Card.Proto.Name, "not enough elements"))
		return
	}

	targetID, _ := json["target"].(string)
	if !checktarget.IsValid(game, seat, power.Target, targetID) {
		log.Error("targetid: ", json["target"])
	}

	seat.Karma.Deactivate(power.Costs)
	if power.Costs.Total() > 0 {
		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
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
