package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/update"
)

const cloningpoolID = "cloning-pool"

func init() {
	game.Scripts[cloningpoolID] = CloningPool
}

func CloningPool(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	var token *game.Token
	if len(args) < 1 {
		return nil, ErrNoTarget
	} else if token, err = target.MyPresentBeing(g, s, args[0]); err != nil {
		return
	} else if token == nil {
		return nil, ErrBadTarget
	} else {
		card := game.NewCard(token.Card.Proto)
		card.Username = token.Username
		token, events = trigger.Spawn(g, s, card)
		token.Body.Health = 1
		if e := trigger.DamageSeat(g, card, s, 1); e != nil {
			events = append(events, e...)
		}

		update.Seat(g, s)
		update.Token(g, token)
	}
	return
}
