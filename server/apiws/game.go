package apiws

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func Game(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user == nil {
			log.Warn("no session")
			socket.Write(websocket.MessageText, out.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", user.Session())

		account := server.Accounts().Get(user.Session().Name())
		if account == nil {
			log.Error("account missing")
			return
		} else if account.GameID == "" {
			log.Warn("no game found")
			socket.Write(websocket.MessageText, out.Error("vii", "you are not in a game"))
			return
		}
		log.Add("Game", account.GameID)

		uri, _ := m.Data["uri"].(string)
		if len(uri) < 1 {
			log.Add("Data", m.Data).Warn("uri missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}
		log.Add("URI", uri)

		if g := server.Games().Get(account.GameID); g == nil {
			log.Warn("game missing")
		} else {
			server.Sessions().Update(account.SessionID)
			g.AddRequest(game.NewReq(account.Username, uri, m.Data))
		}
	})
}
