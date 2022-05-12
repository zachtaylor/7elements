package scripts

import "github.com/zachtaylor/7elements/game/v2"

const vinespiritID = "vine-spirit"

func init() { game.Scripts[vinespiritID] = VineSpirit }

func VineSpirit(g *game.G, ctx game.ScriptContext) ([]game.Phaser, error) {
	if token := g.Token(ctx.Source); token == nil {
		return nil, ErrTokenID
	} else {
		token.T.Body.Attack++
		g.MarkUpdate(token.ID())
		return nil, nil
	}
}
