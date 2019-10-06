package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func Game(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		gameid := m.Data.GetS("gameid")
		uri := m.Data.GetS("uri")
		log := rt.Root.Logger.New().Tag("api/game").With(log.Fields{
			"Username": m.User,
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
			socket.Session.UpdateTime()
			g.Request(m.User, uri, m.Data)
		}
	})
}
