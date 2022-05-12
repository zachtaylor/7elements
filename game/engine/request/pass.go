package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func Pass(game *game.T, seat *seat.T, json map[string]any) {
	log := game.Log().With(map[string]any{
		"State":    game.State,
		"Username": seat.Username,
	})
	if pass, _ := json["pass"].(string); pass == "" {
		log.Warn("target missing")
	} else if pass != game.State.ID() {
		log.Add("PassID", pass).Warn("target mismatch")
	} else if len(game.State.Reacts[seat.Username]) > 0 {
		seat.Writer.WriteMessageData(wsout.Error("pass", "response already recorded"))
	} else {
		game.State.Reacts[seat.Username] = "pass"
		game.Seats.WriteMessageData(wsout.GameReact(game.State.ID(), seat.Username, "pass", game.State.Timer))
	}
}
