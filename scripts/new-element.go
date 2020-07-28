package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func init() {
	game.Scripts["new-element"] = NewElement
}

func NewElement(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if me, ok := me.(*card.T); !ok {
		err = ErrMeCard
	} else {
		events = []game.Stater{state.NewChoice(
			s.Username,
			"Create a New Element",
			cast.JSON{
				"card": me.JSON(),
			},
			out.ChoicesElements,
			func(val interface{}) {
				if i := cast.Int(val); i < 1 || i > 7 {
					out.Error(s.Player, "New Element", "invalid element: "+cast.EscapeString(cast.String(val)))
				} else {
					s.Karma.Append(element.T(i), false)
					out.GameSeat(g, s.JSON())
				}
			},
		)}
	}
	return
}
