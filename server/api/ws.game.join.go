package api

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/http/ws"
	"ztaylor.me/log"
)

func WSGameJoin() ws.Handler {
	return ws.HandlerFunc(func(socket *ws.Socket, m *ws.Message) {
		gameid := m.Data.Sval("id")

		if gameid == "" {
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("game.join: id missing")
			return
		}

		game := game.Service.Get(gameid)

		if game == nil {
			log.WithFields(log.Fields{
				"User":   m.User,
				"GameID": gameid,
			}).Warn("game.join: game missing")
			return
		}

		seat := game.GetSeat(m.User)

		if seat == nil {
			log.WithFields(log.Fields{
				"User":   m.User,
				"GameID": gameid,
			}).Warn("game.join: seat missing")
			return
		}

		seat.Login(game, socket)
		log.WithFields(log.Fields{
			"User":   m.User,
			"GameID": gameid,
		}).Info("game.join")
		<-socket.Done()
		seat.Logout(game)
		log.WithFields(log.Fields{
			"User":   m.User,
			"GameID": gameid,
		}).Debug("game.join: left")
	})
}
