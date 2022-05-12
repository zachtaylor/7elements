package ai

import "github.com/zachtaylor/7elements/game/v2"

type RequestFunc = func(string, map[string]any)

func NewRequestFunc(g *game.G, username string) RequestFunc {
	return func(uri string, json map[string]any) {
		g.Request(username, uri, json)
	}
}
