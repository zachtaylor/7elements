package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func Game(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if socket.Session == nil {
			return
		}
		gameid := m.Data.GetS("gameid")
		uri := m.Data.GetS("uri")
		log := rt.Log().Tag("api/game").With(cast.JSON{
			"Username": socket.Session.Name(),
			"GameID":   gameid,
			"URI":      uri,
		})
		if gameid == "" {
			log.Warn("gameid missing")
		} else if g := rt.Games.Get(gameid); g == nil {
			log.Warn("game missing")
		} else if uri == "" {
			log.Warn("uri missing")
		} else {
			socket.Session.Refresh()
			g.Request(socket.Session.Name(), uri, m.Data)
		}
	})
}
