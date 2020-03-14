package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/update"
)

const summonersportalID = "summoners-portal"

func init() {
	game.Scripts[summonersportalID] = SummonersPortal
}

func SummonersPortal(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	c := s.Deck.Draw()
	if c == nil {
		err = ErrFutureEmpty
	} else if c.Card.Type == card.BodyType || c.Card.Type == card.ItemType {
		if _, _events := trigger.Spawn(g, s, c); len(_events) > 0 {
			events = append(events, _events...)
		}
	} else {
		update.ErrorW(g, "Summoners Portal", "Next card was "+c.Card.Name)
		update.Seat(g, s)
	}
	s.Past[c.ID] = c
	return
}
