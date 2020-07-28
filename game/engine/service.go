package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/power"
	"ztaylor.me/cast"
)

// Service runs this package as a plugable `game.Engine`
type Service struct {
}

// New returns a new `game.Engine`
func New() game.Engine { return &Service{} }

func (s *Service) Run(game *game.T) {
	Run(game)
}

func (s *Service) Start(seat string) game.Stater {
	return state.NewStart(seat)
}

func (s *Service) End(winner, loser string) game.Stater {
	return state.NewEnd(winner, loser)
}

func (s *Service) Target(seat *game.Seat, target string, text string, finish func(val string) []game.Stater) game.Stater {
	return state.NewTarget(seat.Username, target, text, finish)
}

func (s *Service) TriggerTokenEvent(g *game.T, seat *game.Seat, token *game.Token, name string) []game.Stater {
	powers := token.Powers.GetTrigger(name)
	if len(powers) < 1 {
		return nil
	}
	events := make([]game.Stater, 0)
	for _, p := range powers {
		if p.Target != "self" {
			power := p
			events = append(events, state.NewTarget(
				seat.Username,
				p.Target,
				p.Text,
				func(val string) []game.Stater {
					return trigger.Power(g, seat, power, token, cast.NewArray(val))
				},
			))
		} else {
			events = append(events, state.NewTrigger(seat.Username, token, p, token))
		}
	}
	return events
}

func (s *Service) TriggerTokenPower(g *game.T, seat *game.Seat, token *game.Token, power *power.T, arg interface{}) []game.Stater {
	dirty := false
	if power.Costs.Total() > 0 {
		dirty = true
		seat.Karma.Deactivate(power.Costs)
	}
	if power.UsesTurn {
		token.IsAwake = false
		out.GameToken(g, token.JSON())
	}
	if power.UsesKill {
		dirty = true
		delete(seat.Present, token.ID)
	}
	if dirty {
		out.GameSeat(g, seat.JSON())
	}

	if power.Target == "self" {
		return []game.Stater{state.NewTrigger(seat.Username, token, power, token)}
	}
	return []game.Stater{state.NewTrigger(seat.Username, token, power, arg)}
}
