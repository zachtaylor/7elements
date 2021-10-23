package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

// Attack causes phase.Attack to stack
func Attack(game *game.T, seat *seat.T, json websocket.MsgData) (rs []game.Phaser) {
	log := game.Log().With(websocket.MsgData{
		"Seat": seat.String(),
	})

	if id, _ := json["id"].(string); id == "" {
		log.Error("id missing")
	} else if token := seat.Present[id]; token == nil {
		log.Add("ID", id).Error("id invalid")
		seat.Writer.Write(wsout.ErrorJSON(id, "not in your present"))
	} else if token.Body == nil {
		log.Add("Token", token.String()).Error("card type must be body")
		seat.Writer.Write(wsout.ErrorJSON(token.Card.Proto.Name, `not "body" type`))
	} else if !token.IsAwake {
		log.Add("Token", token.String()).Error("card must be awake")
		seat.Writer.Write(wsout.ErrorJSON(token.Card.Proto.Name, "not awake"))
	} else {
		log.Add("Token", token.String()).Info("accept")
		rs = append(rs, game.Engine().SleepToken(game, token)...)
		rs = append(rs, phase.NewAttack(seat.Username, token))
	}
	return
}
