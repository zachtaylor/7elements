package request

import (
	"github.com/zachtaylor/7elements/game"
	pkg_state "github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

// attack causes AttackEvent to stack on MainEvent
func attack(g *game.T, seat *game.Seat, json cast.JSON) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.String(),
	}).Tag("engine/attack")

	if id := json.GetS("id"); id == "" {
		log.Error("id missing")
	} else if token := seat.Present[id]; token == nil {
		log.Add("ID", id).Error("id invalid")
		update.ErrorW(seat, id, `not in your present`)
	} else if token.Body == nil {
		log.Add("Token", token.String()).Error("card type must be body")
		update.ErrorW(seat, token.Card.Card.Name, `not "body" type`)
	} else if !token.IsAwake {
		log.Add("Token", token.String()).Error("card must be awake")
		update.ErrorW(seat, token.Card.Card.Name, `not awake`)
	} else {
		log.Add("Token", token.String()).Info("accept")
		token.IsAwake = false
		update.Token(g, token)
		return []game.Stater{pkg_state.NewAttack(seat.Username, token)}
	}
	return nil
}
