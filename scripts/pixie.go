package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

const PixieID = "pixie"

func init() {
	game.Scripts[PixieID] = Pixie
}

func Pixie(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = game.ErrNoTarget
	} else if token, ok := me.(*game.Token); !ok {
		err = game.ErrMeToken
	} else if token == nil {
		err = game.ErrMeNil
	} else if seat := g.Seats[cast.String(args[0])]; seat != nil {
		hp := token.Body.Health
		if _events := trigger.HealSeat(g, token.Card, seat, hp); len(_events) > 0 {
			events = append(events, _events...)
		}
	} else if targetToken, _err := target.OtherPresentBeing(g, s, token, args[0]); _err != nil {
		err = _err
	} else if targetToken == nil {
		err = game.ErrBadTarget
	} else {
		hp := token.Body.Health
		events = append(events, trigger.Death(g, token)...)
		events = append(events, state.NewTarget(
			s.Username,
			"player-being",
			cast.StringN("Target Player or Being gains", hp, "Health"),
			func(val string) []game.Stater {
				if val == s.Username {
					return trigger.HealSeat(g, token.Card, s, hp)
				} else if opponent := g.GetOpponentSeat(s.Username); val == opponent.Username {
					return trigger.HealSeat(g, token.Card, opponent, hp)
				}
				token, err := target.PresentBeing(g, s, val)
				if err != nil {
					g.Log().Add("Error", err).Error("scripts/pixie")
					return nil
				}
				token.Body.Health += hp
				update.Token(g, token)
				return nil
			},
		))
	}
	return
}
