package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func Login(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user != nil {
			log.Add("Session", user.Session()).Warn("session exists")
			socket.Write(websocket.MessageText, out.Error("vii", "you are already logged in!"))
			return
		}

		username, _ := m.Data["username"].(string)
		if username == "" {
			log.Warn("username missing")
			socket.Write(websocket.MessageText, out.Error("vii", "username missing"))
			return
		} else if err := api.CheckUsername(username); err != nil {
			log.With(map[string]any{
				"Name":  username,
				"Error": err.Error(),
			}).Warn("username not allowed")
			socket.Write(websocket.MessageText, out.Error("vii", "username banned"))
			return
		}
		log = log.Add("Name", username)

		var password string
		if passbuff, _ := m.Data["password"].(string); len(passbuff) > 0 {
			password = server.HashPassword(passbuff)
		} else {
			log.Warn("password missing")
			socket.Write(websocket.MessageText, out.Error("vii", "password missing"))
			return
		}

		account, session, err := api.Login(server, username, password)
		if account == nil || session == nil {
			log.Add("Error", err).Error("failed")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}
		log = log.Add("Session", session.ID())

		log.Info("ok")
		server.Users().Must(socket, username)
		socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
		socket.WriteMessage(websocket.NewMessage("/redirect", map[string]any{
			"location": "/",
		}))
	})
}
