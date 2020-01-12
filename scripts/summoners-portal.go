package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game/trigger"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

const summonersportalID = "summoners-portal"

func init() {
	game.Scripts[summonersportalID] = SummonersPortal
}

func SummonersPortal(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	card := s.Deck.Draw()
	if card == nil {
		err = game.ErrFutureEmpty
	} else if card.Card.Type == vii.CTYPbody || card.Card.Type == vii.CTYPitem {
		if _, _events := trigger.Spawn(g, s, card); len(_events) > 0 {
			events = append(events, _events...)
		}
	} else {
		update.ErrorW(g, "Summoners Portal", "Next card was "+card.Card.Name)
		update.Seat(g, s)
	}
	s.Past[card.ID] = card
	return
}
