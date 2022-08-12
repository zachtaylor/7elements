package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func GameCancel(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user == nil {
			log.Warn("no session")
			socket.Write(websocket.MessageText, out.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", user.Session())

		if !server.MatchMaker().Cancel(user.Session().Name()) {
			log.Warn("failed")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
		} else {
			log.Info("ok")
			socket.WriteMessage(websocket.NewMessage("/game/queue", nil))
		}
	})
}
