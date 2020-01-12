package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

const SymbiosisID = "symbiosis"

func init() {
	game.Scripts[SymbiosisID] = Symbiosis
}

func Symbiosis(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	log := g.Log().With(cast.JSON{
		"Username": s.Username,
	}).Tag(logtag + SymbiosisID)

	return []game.Stater{state.NewTarget(
		s.Username,
		"being",
		"Target Being gains 1 Attack",
		func(val string) []game.Stater {
			token, err := target.PresentBeing(g, s, val)
			if err != nil {
				log.Add("Error", err).Error()
			} else {
				log.Add("Token", token.String()).Info()
				token.Body.Attack++
				update.Token(g, token)
			}
			return nil
		},
	)}, nil
}
