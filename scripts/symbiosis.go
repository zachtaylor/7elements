package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/out"
)

const SymbiosisID = "symbiosis"

func init() {
	game.Scripts[SymbiosisID] = Symbiosis
}

func Symbiosis(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	return []game.Stater{state.NewTarget(
		s.Username,
		"being",
		"Target Being gains 1 Attack",
		func(val string) []game.Stater {
			token, err := target.PresentBeing(g, s, val)
			if err != nil {
				g.Log().Add("Error", err).Error()
			} else {
				g.Log().Add("Token", token.String()).Info()
				token.Body.Attack++
				out.GameToken(g, token.JSON())
			}
			return nil
		},
	)}, nil
}
