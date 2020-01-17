package scripts

import (
	"github.com/zachtaylor/7elements/game/update"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/cast"
)

const bendwillID = "bend-will"

func init() {
	game.Scripts[bendwillID] = BendWill
}

func BendWill(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, e := target.PresentBeing(g, s, args[0]); err != nil {
		err = e
	} else {
		events = []game.Stater{state.NewChoice(
			s.Username,
			cast.StringN(token.Card.Card.Name, "-> Awake or Asleep?"),
			cast.JSON{
				"token": token.JSON(),
			},
			[]cast.JSON{
				cast.JSON{"choice": "awake", "display": `<a>Awake</a>`},
				cast.JSON{"choice": "asleep", "display": `<a>Asleep</a>`},
			},
			func(val interface{}) {
				if val == "awake" {
					token.IsAwake = true
				} else if val == "asleep" {
					token.IsAwake = false
				} else {
					return
				}
				update.Token(g, token)
			},
		)}
	}
	return
}
