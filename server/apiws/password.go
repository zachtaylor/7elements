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
	log := rt.Logger.Add("Socket", socket.ID())
	if len(socket.Name()) < 1 {
		log.Warn("anon update email")
		socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to change email"))
		return
	}
	log.Add("Name", socket.Name())

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

	account := rt.Accounts.Get(socket.Name())
	if account == nil {
		log.Warn("password mismatch")
		socket.WriteSync(wsout.ErrorJSON("vii", "password2 missing"))
		return
	}
	log.Info()

	account.Password = pass2
	if err := accounts.UpdatePassword(rt.DB, account); err != nil {
		log.Add("Error", err.Error()).Error("update password")
		socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
		return
	}
	socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
}
