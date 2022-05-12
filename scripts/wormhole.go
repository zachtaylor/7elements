package scripts

import (
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/game/v2/target"
)

const WormholeID = "wormhole"

func init() { game.Scripts[WormholeID] = Wormhole }

func Wormhole(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if token, err := target.PresentBeing(g, ctx.Targets[0]); err != nil {
		return nil, err
	} else if tokenOwner := g.Player(token.Player()); tokenOwner == nil {
		return nil, ErrPlayerID
	} else if card := g.Card(token.T.Card); card == nil {
		return nil, ErrCardID
	} else if cardOwner := g.Player(card.Player()); cardOwner == nil {
		return nil, ErrPlayerID
	} else {
		g.Log().With(map[string]any{
			"Target":  token.ID(),
			"T Owner": tokenOwner.ID(),
			"C Owner": cardOwner.ID(),
		}).Info()
		rs := trigger.TokenRemove(g, token)

		if cardOwner.T.Past.Has(card.ID()) {
			delete(cardOwner.T.Past, card.ID())
			cardOwner.T.Hand.Set(card.ID())
			g.MarkUpdate(cardOwner.ID())
			out.PrivateCard(g, cardOwner, card.ID())
		}
		return rs, nil
	}
}
