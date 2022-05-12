package apiws

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Password(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		password(server, socket, m)
	})
}

func password(server internal.Server, socket *websocket.T, m *websocket.Message) {
	log := server.Log().Add("Socket", socket.ID())

	if len(socket.SessionID()) < 1 {
		log.Warn("no session")
		socket.Write(wsout.Error("vii", "no user"))
		return
	}
	log = log.Add("Session", socket.SessionID())

	user, _, err := server.GetUserManager().GetSession(socket.SessionID())
	if user == nil {
		log.Add("Error", err).Error("user missing")
		socket.Write(wsout.Error("vii", "internal error"))
		return
	}
	log = log.Add("Username", user.Name())

	account := server.GetAccounts().Get(user.Name())
	if account == nil {
		log.Add("User", user.Name()).Error("no account")
		socket.Write(wsout.Error("vii", "no account"))
		return
	}
	log.Info()

	passbuff1, _ := m.Data["password1"].(string)
	if passbuff1 == "" {
		log.Warn("password1 missing")
		socket.WriteSync(wsout.Error("vii", "password1 missing"))
		return
	}
	pass1 := server.HashPassword(passbuff1)

	passbuff2, _ := m.Data["password2"].(string)
	if passbuff2 == "" {
		log.Warn("password2 missing")
		socket.WriteSync(wsout.Error("vii", "password2 missing"))
		return
	}
	pass2 := server.HashPassword(passbuff2)

	if pass1 != pass2 {
		log.Warn("password mismatch")
		socket.WriteSync(wsout.Error("vii", "password2 missing"))
		return
	}

	if err := accounts.UpdatePassword(server.GetDB(), account); err != nil {
		log.Add("Error", err.Error()).Error("update password")
		socket.WriteSync(wsout.Error("vii", "internal error"))
		return
	}
	account.Password = pass2

	log.Info()
	socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
}
