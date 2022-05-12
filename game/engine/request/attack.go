package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

// Attack causes phase.Attack to stack
func Attack(game *game.T, seat *seat.T, json map[string]any) (rs []game.Phaser) {
	log := game.Log().With(map[string]any{
		"Seat": seat.String(),
	})

	if id, _ := json["id"].(string); id == "" {
		log.Error("id missing")
	} else if token := seat.Present[id]; token == nil {
		log.Add("ID", id).Error("id invalid")
		seat.Writer.WriteMessageData(wsout.Error(id, "not in your present"))
	} else if token.Body == nil {
		log.Add("Token", token.String()).Error("card type must be body")
		seat.Writer.WriteMessageData(wsout.Error(token.Card.Proto.Name, `not "body" type`))
	} else if !token.IsAwake {
		log.Add("Token", token.String()).Error("card must be awake")
		seat.Writer.WriteMessageData(wsout.Error(token.Card.Proto.Name, "not awake"))
	} else {
		log.Add("Token", token.String()).Info("accept")
		rs = append(rs, game.Engine().SleepToken(game, token)...)
		rs = append(rs, phase.NewAttack(seat.Username, token))
	}
	return
}
