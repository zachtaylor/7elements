package trigger

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func Spawn(g *game.T, seat *game.Seat, card *card.T) (*game.Token, []game.Stater) {
	token := game.NewToken(card, seat.Username)
	g.RegisterToken(token)
	seat.Present[token.ID] = token
	update.Token(g, token)
	return token, nil // TODO
}
