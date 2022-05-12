package apiws

import (
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func GameCancel(server internal.Server) websocket.Handler {
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

		if err := server.GetMatchMaker().Cancel(user.Name()); err != nil {
			log.Add("Error", err).Warn("failed")
			socket.Write(wsout.Error("vii", "internal error"))
		} else {
			log.Info("ok")
			socket.WriteSync(wsout.Queue(nil))
		}
	})
}
