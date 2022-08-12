package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/target"
)

const memorializeID = "memorialize"

func init() { game.Scripts[memorializeID] = Memorialize }

func Memorialize(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if len(ctx.Targets) < 1 {
		return nil, ErrNoTarget
	} else if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else if targetCard, err := target.MyPastBeing(g, player, ctx.Targets[0]); err != nil {
		return nil, err
	} else {
		card := g.NewCard(player.ID(), targetCard.T)
		player.T.Hand.Add(card.ID())
		out.PrivateCard(player, card)
		out.PrivateHand(player)
		return nil, nil
	}
}
