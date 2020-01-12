package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

const crystalballID = "crystal-ball"

func init() {
	game.Scripts[crystalballID] = CrystalBall
}

func CrystalBall(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = game.ErrNoTarget
	} else {
		card := s.Deck.Cards[0]
		events = []game.Stater{state.NewChoice(
			s.Username,
			"Shuffle your Future?",
			cast.JSON{
				"card": card.JSON(),
			},
			[]cast.JSON{
				cast.JSON{"choice": true, "display": `<a>Yes</a>`},
				cast.JSON{"choice": false, "display": `<a>No</a>`},
			},
			func(val interface{}) {
				if cast.Bool(val) {
					s.Deck.Shuffle()
					update.Seat(g, s)
				} else {
				}
			},
		)}
	}
	return
}
