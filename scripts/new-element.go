package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func init() {
	game.Scripts["new-element"] = NewElement
}

func NewElement(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if me, ok := me.(*game.Card); !ok {
		err = game.ErrMeCard
	} else {
		events = []game.Stater{state.NewChoice(
			s.Username,
			"Create a New Element",
			cast.JSON{
				"card": me.JSON(),
			},
			update.ChoicesNewElement,
			func(val interface{}) {
				if i := cast.Int(val); i < 1 || i > 7 {
					update.ErrorW(g, "New Element", "invalid element: "+cast.EscapeString(cast.String(val)))
				} else {
					e := vii.Element(i)
					s.Elements.Append(e)
					update.Seat(g, s)
				}
			},
		)}
	}
	return
}
