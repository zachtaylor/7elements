package request

import (
	"github.com/zachtaylor/7elements/game"
	pkg_state "github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

// attack causes AttackEvent to stack on MainEvent
func attack(g *game.T, seat *game.Seat, json cast.JSON) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.String(),
	})

	if id := json.GetS("id"); id == "" {
		log.Error("id missing")
	} else if token := seat.Present[id]; token == nil {
		log.Add("ID", id).Error("id invalid")
		out.Error(seat.Player, id, `not in your present`)
	} else if token.Body == nil {
		log.Add("Token", token.String()).Error("card type must be body")
		out.Error(seat.Player, token.Card.Proto.Name, `not "body" type`)
	} else if !token.IsAwake {
		log.Add("Token", token.String()).Error("card must be awake")
		out.Error(seat.Player, token.Card.Proto.Name, `not awake`)
	} else {
		log.Add("Token", token.String()).Info("accept")
		token.IsAwake = false
		out.GameToken(g, token.JSON())
		return []game.Stater{pkg_state.NewAttack(seat.Username, token)}
	}
	return nil
}
