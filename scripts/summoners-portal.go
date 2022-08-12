package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"taylz.io/yas"
)

const summonersportalID = "summoners-portal"

func init() { game.Scripts[summonersportalID] = SummonersPortal }

func SummonersPortal(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	player := g.Player(ctx.Player)
	if player == nil {
		return nil, ErrPlayerID
	} else if len(player.T.Future) < 1 {
		return nil, ErrFutureEmpty
	}
	cardID, future := yas.Shift(player.T.Future)
	player.T.Future = future
	c := g.Card(cardID)
	g.MarkUpdate(ctx.Player)
	if c == nil {
		return nil, ErrCardID
	} else if c.T.Kind == card.Being || c.T.Kind == card.Item {
		return []game.Phaser{
			phase.NewPlay(g, ctx.Player, c, element.Count{}, []string{}),
		}, nil
	} else {
		player.T.Past.Add(cardID)
		// seat.Writer.WriteMessageData(wsout.Error("Summoners Portal", "Next card was "+c.Proto.Name))
		return nil, nil
	}
}
