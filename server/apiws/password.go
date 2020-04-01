package apiws

import (
	"github.com/zachtaylor/7elements/game/update"
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/http/websocket"
)

func Password(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		password(rt, socket, m)
	})
}

func password(rt *Runtime, socket *websocket.T, m *websocket.Message) {
	log := rt.Runtime.Root.Logger.New().Add("Socket", socket).Add("Message", m)
	log.Copy().Source().Info()
	session := socket.Session
	if session == nil {
		update.ErrorSock(socket, "password change", "no session")
	} else if account := rt.Runtime.Root.Accounts.Test(session.Name()); account == nil {
		update.ErrorSock(socket, "password change", "no account")
	} else if newp1, newp2 := internal.HashPassword(m.Data.GetS("password1"), rt.Runtime.Salt), internal.HashPassword(m.Data.GetS("password2"), rt.Runtime.Salt); newp1 != newp2 {
		update.ErrorSock(socket, "password change", "password mismatch")
	} else {
		log.Copy().Source().Trace(`about to engage`)
		account.Password = newp2
		if err := rt.Runtime.Root.Accounts.UpdatePassword(account); err != nil {
			update.ErrorSock(socket, "password change", err.Error())
		}
		socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(account))
	}
}
