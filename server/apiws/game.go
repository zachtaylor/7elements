package apiws

import (
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Game(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("no session")
			socket.Write(wsout.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", socket.SessionID())

		user, _, err := server.GetUserManager().GetSession(socket.SessionID())
		if user == nil {
			log.Add("Error", err).Error("user missing")
			socket.Write(wsout.Error("vii", "internal error"))
			return
		}
		log = log.Add("Username", user.Name())

		account := server.GetAccounts().Get(user.Name())
		if account == nil {
			log.Error("account missing")
			return
		} else if account.GameID == "" {
			log.Warn("no game found")
			socket.WriteSync(wsout.Error("vii", "you are not in a game"))
			return
		}
		log.Add("Game", account.GameID)

		uri, _ := m.Data["uri"].(string)
		if len(uri) < 1 {
			log.Add("Data", m.Data).Warn("uri missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		}
		log.Add("URI", uri)

		if game := server.GetGameManager().Get(account.GameID); game == nil {
			log.Warn("game missing")
		} else {
			server.GetSessionManager().Get(account.SessionID).Update()
			game.Request(user.Name(), uri, m.Data)
		}
	})
}
