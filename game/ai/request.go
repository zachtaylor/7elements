package ai

import (
	"github.com/zachtaylor/7elements/game"
	"taylz.io/http/websocket"
)

type RequestFunc = func(string, websocket.MsgData)

func NewRequestFunc(g *game.T, username string) RequestFunc {
	return func(uri string, json websocket.MsgData) {
		g.Request(username, uri, json)
	}
}
