package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Login(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())
		if len(socket.SessionID()) > 0 {
			log.Add("Session", socket.SessionID()).Warn("session exists")
			socket.WriteSync(wsout.Error("vii", "you are already logged in!"))
			return
		}

		username, _ := m.Data["username"].(string)
		if username == "" {
			log.Warn("username missing")
			socket.WriteSync(wsout.Error("vii", "username missing"))
			return
		} else if err := api.CheckUsername(username); err != nil {
			log.With(map[string]any{
				"Name":  username,
				"Error": err.Error(),
			}).Warn("username not allowed")
			socket.WriteSync(wsout.Error("vii", "username banned"))
			return
		}
		log = log.Add("Name", username)

		var password string
		if passbuff, _ := m.Data["password"].(string); len(passbuff) > 0 {
			password = server.HashPassword(passbuff)
		} else {
			log.Warn("password missing")
			socket.WriteSync(wsout.Error("vii", "password missing"))
			return
		}

		account, session, err := api.Login(server, username, password)
		if account == nil || session == nil {
			log.Add("Error", err).Error("failed")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		}
		log = log.Add("Session", session.ID())

		log.Info("ok")
		server.GetUserManager().Authorize(username, socket)
		socket.Write(wsout.MyAccount(account.Data()).EncodeToJSON())
		socket.Write(wsout.Redirect("/").EncodeToJSON())
	})
}
