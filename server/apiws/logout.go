package apiws

import (
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Logout(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("anon logout")
			socket.WriteSync(wsout.Error("vii", "you must log in to log out"))
			return
		}

		session := server.GetSessionManager().Get(socket.SessionID())
		if session == nil {
			log.Add("Session", socket.SessionID()).Warn("session not found")
			socket.WriteSync(wsout.Error("vii", "you must log in to log out"))
			return
		}

		if server.GetMatchMaker().Get(session.Name()) != nil {
			log.Debug("cancel queue")
			server.GetMatchMaker().Cancel(session.Name())
			socket.WriteSync(wsout.Queue(nil))
		}

		log.Info("ok")
		server.GetSessionManager().Remove(session.ID())
		socket.WriteSync(wsout.MyAccount(nil).EncodeToJSON())
	})
}
