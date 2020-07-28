package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/runtime"
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/http/websocket"
)

func Password(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		password(rt, socket, m)
	})
}

func password(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
	log := rt.Log().Add("Socket", socket).Add("Message", m)
	log.Info()
	newp1 := internal.HashPassword(m.Data.GetS("password1"), rt.PassSalt)
	newp2 := internal.HashPassword(m.Data.GetS("password2"), rt.PassSalt)

	if socket.Session == nil {
		out.Error(socket, "password change", "no session")
	} else if account, _ := rt.Accounts.Get(socket.Session.Name()); account == nil {
		out.Error(socket, "password change", "no account")
	} else if newp1 != newp2 {
		out.Error(socket, "password change", "password mismatch")
	} else {
		log.Copy().Trace(`about to engage`)
		account.Password = newp2
		if err := rt.Accounts.UpdatePassword(account); err != nil {
			out.Error(socket, "password change", err.Error())
		}
		socket.Send("/myaccount", rt.AccountJSON(account))
	}
}
