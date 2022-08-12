package request

import (
	"reflect"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

func Trigger(g *game.G, state *game.State, player *game.Player, json map[string]any) (rs []game.Phaser) {
	log := g.Log().Add("Player", player.ID())

	// validation

	tokenID, _ := json["id"].(string)
	if len(tokenID) < 1 {
		log.Error("id: ", json["id"])
		return
	}

	log = log.Add("TokenID", tokenID)

	token, err := target.MyPresent(g, player, tokenID)
	if err != nil {
		log.Add("Error", err).Error()
		player.T.Writer.Write(out.ErrorMessage("trigger", err.Error()))
		return nil
	}

	var powerID int
	if poweridbuff, _ := json["powerid"].(float64); int(poweridbuff) < 1 {
		g.Log().Add("powerid", json["powerid"]).Add("type", reflect.TypeOf(json["powerid"])).Warn("powerid missing")
		return nil
	} else {
		powerID = int(poweridbuff)
	}

	log = log.Add("PowerId", powerID)

	power := token.T.Powers[powerID]
	if power == nil {
		log.Add("Keys", token.T.Powers.Keys()).Error("power not found")
		return
	} else if !token.T.Awake && power.UsesTurn {
		log.Error("card is asleep")
		player.T.Writer.Write(out.ErrorMessage(token.T.Name, "not awake"))
		return
	} else if !player.T.Karma.Active().Test(power.Costs) {
		log.Add("Costs", power.Costs).Error("cannot afford")
		player.T.Writer.Write(out.ErrorMessage(token.T.Name, "not enough elements"))
		return
	}

	targetID, _ := json["target"].(string)
	if !target.IsValid(g, player, power.Target, targetID) {
		log.Error("targetid: ", json["target"])
	}

	player.T.Karma.Deactivate(power.Costs)
	if power.Costs.Total() > 0 {
		g.MarkUpdate(player.ID())
	}
	if power.UsesTurn {
		if triggered := trigger.TokenSleep(g, token); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}
	if power.UsesLife {
		if triggered := trigger.TokenRemove(g, token); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}

	log.Add("TargetID", targetID).Add("Power", power).Trace()

	rs = append(rs, game.NewTriggerPhase(g, token, power))
	return
}
