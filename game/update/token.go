package update

import "github.com/zachtaylor/7elements/game"

func Token(g *game.T, t *game.Token) {
	g.Log().Tag("update/token").Add("Owner", t.Username).Debug()
	Game(g, "/game/token", t.JSON())
}
