package apiws

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Password(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		password(rt, socket, m)
	})
}

func password(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
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

	account := rt.Accounts.Get(user.Name())
	if account == nil {
		log.Add("User", user.Name()).Error("no account")
		socket.Write(wsout.ErrorJSON("vii", "no account"))
		return
	}
	log.Info()

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

	if err := accounts.UpdatePassword(rt.DB, account); err != nil {
		log.Add("Error", err.Error()).Error("update password")
		socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
		return
	}
	account.Password = pass2

	log.Info()
	socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
}
