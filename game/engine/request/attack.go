package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/trigger"
)

// Attack causes phase.Attack to stack
func Attack(game *game.G, state *game.State, player *game.Player, json map[string]any) (rs []game.Phaser) {
	log := game.Log().Add("Player", player.T.Writer.Name())

	if state.T.Phase.Type() != "main" {
		player.T.Writer.Write(out.ErrorMessage("vii", "not in your main"))
	} else if state.T.Phase.Priority()[0] != player.ID() {
		player.T.Writer.Write(out.ErrorMessage("vii", "not your turn"))
	}

	if id, _ := json["id"].(string); id == "" {
		log.Error("id missing")
	} else if _, ok := player.T.Present[id]; ok {
		log.Add("ID", id).Error("id invalid")
		player.T.Writer.Write(out.ErrorMessage(id, "not in your present"))
	} else if token := game.Token(id); token == nil {
		log.Add("ID", id).Error("token error")
		player.T.Writer.Write(out.ErrorMessage(id, "tokenid error"))
	} else if token.T.Body == nil {
		log.Add("Token", token.ID()).Error("card type must be body")
		player.T.Writer.Write(out.ErrorMessage(token.T.Name, `not "body" type`))
	} else if !token.T.Awake {
		log.Add("Token", token.ID()).Error("card must be awake")
		player.T.Writer.Write(out.ErrorMessage(token.T.Name, "not awake"))
	} else {
		log.Add("Token", token.ID()).Info("accept")
		rs = append(rs, trigger.TokenSleep(game, token)...)
		rs = append(rs, phase.NewAttack(game, token))
	}
	return
}
