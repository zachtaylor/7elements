package api

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/http/ws"
	"ztaylor.me/log"
)

func WSGame() ws.Handler {
	return ws.HandlerFunc(func(socket *ws.Socket, m *ws.Message) {
		gameid := m.Data.Sval("gameid")

		if gameid == "" {
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("api/ws/game: gameid missing")
			return
		}

		g := game.Service.Get(gameid)

		if g == nil {
			log.WithFields(log.Fields{
				"User":   m.User,
				"GameID": gameid,
			}).Warn("api/ws/game: game missing")
			return
		}

		uri := m.Data.Sval("uri")

		if uri == "" {
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("api/ws/game: stateid missing")
			return
		}

		g.Request(m.User, uri, m.Data)
	})
}
