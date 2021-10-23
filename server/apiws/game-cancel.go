package apiws

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func GameCancel(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("no session")
			socket.Write(wsout.ErrorJSON("vii", "no user"))
			return
		}
		log = log.Add("Session", socket.SessionID())

		user, _, err := rt.Users.GetSession(socket.SessionID())
		if user == nil {
			log.Add("Error", err).Error("user missing")
			socket.Write(wsout.ErrorJSON("vii", "internal error"))
			return
		}
		log = log.Add("Username", user.Name())

		if err := rt.MatchMaker.Cancel(user.Name()); err != nil {
			log.Add("Error", err).Warn("failed")
			socket.Write(wsout.ErrorJSON("vii", "internal error"))
		} else {
			log.Info("ok")
			socket.WriteSync(wsout.Queue(nil))
		}
	})
}
