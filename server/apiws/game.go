package apiws

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Game(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Logger.Add("Socket", socket.ID())
		if len(socket.Name()) < 1 {
			log.Warn("anon game")
			socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to play game"))
			return
		}
		log.Add("Name", socket.Name())

		account := rt.Accounts.Get(socket.Name())
		if account == nil {
			log.Error("account missing")
			return
		} else if account.GameID == "" {
			log.Warn("no game found")
			socket.WriteSync(wsout.ErrorJSON("vii", "you are not in a game"))
			return
		}
		log.Add("Game", account.GameID)

		uri, _ := m.Data["uri"].(string)
		if len(uri) < 1 {
			log.Add("Data", m.Data).Warn("uri missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}
		log.Add("URI", uri)

		if game := rt.Games.Get(account.GameID); game == nil {
			log.Warn("game missing")
		} else {
			rt.Sessions.Get(account.SessionID).Update()
			game.Request(socket.Name(), uri, m.Data)
		}
	})
}
