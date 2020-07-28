package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/out"
)

const summonersportalID = "summoners-portal"

func init() {
	game.Scripts[summonersportalID] = SummonersPortal
}

func SummonersPortal(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	c := s.Deck.Draw()
	if c == nil {
		err = ErrFutureEmpty
	} else if c.Proto.Type == card.BodyType || c.Proto.Type == card.ItemType {
		if _, _events := trigger.Spawn(g, s, c); len(_events) > 0 {
			events = append(events, _events...)
		}
	} else {
		out.Error(s.Player, "Summoners Portal", "Next card was "+c.Proto.Name)
		out.GameSeat(g, s.JSON())
	}
	s.Past[c.ID] = c
	return
}
