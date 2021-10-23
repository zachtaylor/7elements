package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Login(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Logger.Add("Socket", socket.ID()).Add("Data", m.Data)
		if len(socket.SessionID()) > 0 {
			log.Add("Session", socket.SessionID()).Warn("session exists")
			socket.WriteSync(wsout.ErrorJSON("vii", "you are already logged in!"))
			return
		}

		username, _ := m.Data["username"].(string)
		if username == "" {
			log.Warn("username missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "username missing"))
			return
		} else if err := api.CheckUsername(username); err != nil {
			log.With(map[string]interface{}{
				"Name":  username,
				"Error": err.Error(),
			}).Warn("username not allowed")
			socket.WriteSync(wsout.ErrorJSON("vii", "username banned"))
			return
		}
		log = log.Add("Name", username)

		var password string
		if passbuff, _ := m.Data["password"].(string); len(passbuff) > 0 {
			password = rt.PassHash(passbuff)
		} else {
			log.Warn("password missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "password missing"))
			return
		}

		account, session, err := api.Login(rt, username, password)
		if account == nil || session == nil {
			log.Add("Error", err).Error("failed")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		log.Info("ok")
		rt.Users.Authorize(username, socket)
		socket.Write(wsout.MyAccount(account.Data()).EncodeToJSON())
		socket.Write(wsout.Redirect("/").EncodeToJSON())
	})
}
