package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func Logout(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user == nil {
			log.Warn("no session")
			socket.Write(websocket.MessageText, out.Error("vii", "you must log in to log out"))
			return
		}
		log = log.Add("Session", user.Session())

		if server.MatchMaker().Get(user.Session().Name()) == nil {
			log.Debug("not in queue")
		} else {
			log.Debug("cancelling queue...")
			server.MatchMaker().Cancel(user.Session().Name())
			socket.WriteMessage(websocket.NewMessage("/game/queue", nil))
		}

		log.Info("ok")
		server.Sessions().Remove(user.Session().ID())
		socket.WriteMessage(websocket.NewMessage("/myaccount", nil))
	})
}
