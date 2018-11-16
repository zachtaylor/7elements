package api

import (
	"github.com/zachtaylor/7elements"
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
			}).Warn("game.event: id missing")
			return
		}

		game := vii.GameService.Get(gameid)

		if game == nil {
			log.WithFields(log.Fields{
				"User":   m.User,
				"GameID": gameid,
			}).Warn("game.event: game missing")
			return
		}

		game.In <- &vii.GameRequest{
			Username: m.User,
			Data:     m.Data,
		}
	})
}
