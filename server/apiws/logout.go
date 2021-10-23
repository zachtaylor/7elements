package apiws

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Logout(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("anon logout")
			socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to log out"))
			return
		}

		session := rt.Sessions.Get(socket.SessionID())
		if session == nil {
			log.Add("Session", socket.SessionID()).Warn("session not found")
			socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to log out"))
			return
		}

		if rt.MatchMaker.Get(session.Name()) != nil {
			log.Debug("cancel queue")
			rt.MatchMaker.Cancel(session.Name())
			socket.WriteSync(wsout.Queue(nil))
		}

		log.Info("ok")
		rt.Sessions.Remove(session.ID())
		socket.WriteSync(wsout.MyAccount(nil).EncodeToJSON())
	})
}
