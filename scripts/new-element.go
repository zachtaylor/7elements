package scripts

import (
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
)

func init() { game.Scripts["new-element"] = NewElement }

func NewElement(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if player := g.Player(ctx.Player); player == nil {
		return nil, ErrPlayerID
	} else {
		return []game.Phaser{phase.NewChoice(
			ctx.Player,
			"Create a New Element",
			nil,
			phase.ChoiceElementsData,
			func(val any) {
				if i, _ := val.(int); i < 1 {
					g.Log().Error("type", val)
				} else if el := element.T(i); el < element.White || el > element.Black {
					g.Log().Error("invalid element", i)
				} else {
					player.T.Karma.Append(el, false)
					g.MarkUpdate(player.ID())
				}
			},
		)}, nil
	}
}
