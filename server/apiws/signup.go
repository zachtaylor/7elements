package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Signup(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Logger.Add("Socket", socket.ID())
		if len(socket.SessionID()) > 1 {
			log.Add("Session", socket.SessionID()).Warn("signup")
			socket.WriteSync(wsout.ErrorJSON("vii", "you are already logged in"))
			return
		}

		username, _ := m.Data["username"].(string)
		if username == "" {
			log.Warn("username missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "username missing"))
			return
		}
		log = log.Add("Name", username)

		email, _ := m.Data["email"].(string)
		if email == "" {
			log.Add("Data", m.Data).Warn("email missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "no new email"))
			return
		}
		log = log.Add("Email", email)

		passbuff1, _ := m.Data["password1"].(string)
		if passbuff1 == "" {
			log.Warn("password1 missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "password1 missing"))
			return
		}
		pass1 := rt.PassHash(passbuff1)

		passbuff2, _ := m.Data["password2"].(string)
		if passbuff2 == "" {
			log.Warn("password2 missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "password2 missing"))
			return
		}
		pass2 := rt.PassHash(passbuff2)

		if pass1 != pass2 {
			log.Warn("password mismatch")
			socket.WriteSync(wsout.ErrorJSON("vii", "password2 missing"))
			return
		}

		account, session, err := api.Signup(rt, username, email, pass1)
		if account == nil || session == nil {
			log.Add("Error", err).Error("failed")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		log.Add("Session", session.ID()).Info("ok")

		rt.Users.Authorize(username, socket)
		socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
		socket.WriteSync(wsout.Redirect("/").EncodeToJSON())
		rt.Ping()
	})
}
