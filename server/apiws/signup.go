package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func Signup(server internal.Server) websocket.MessageHandler {
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
		}
		log = log.Add("Name", username)

		email, _ := m.Data["email"].(string)
		if email == "" {
			log.Add("Data", m.Data).Warn("email missing")
			socket.Write(websocket.MessageText, out.Error("vii", "no new email"))
			return
		}
		log = log.Add("Email", email)

		passbuff1, _ := m.Data["password1"].(string)
		if passbuff1 == "" {
			log.Warn("password1 missing")
			socket.Write(websocket.MessageText, out.Error("vii", "password1 missing"))
			return
		}
		pass1 := server.HashPassword(passbuff1)

		passbuff2, _ := m.Data["password2"].(string)
		if passbuff2 == "" {
			log.Warn("password2 missing")
			socket.Write(websocket.MessageText, out.Error("vii", "password2 missing"))
			return
		}
		pass2 := server.HashPassword(passbuff2)

		if pass1 != pass2 {
			log.Warn("password mismatch")
			socket.Write(websocket.MessageText, out.Error("vii", "password2 missing"))
			return
		}

		account, session, err := api.Signup(server, username, email, pass1)
		if account == nil || session == nil {
			log.Add("Error", err).Error("failed")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		log.Add("Session", session.ID()).Info("ok")

		server.Users().Must(socket, username)
		socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
		socket.WriteMessage(websocket.NewMessage("/redirect", map[string]any{
			"location": "/",
		}))
	})
}
