package apiws

import (
	"github.com/zachtaylor/7elements/game/update"
	"github.com/zachtaylor/7elements/server/api"
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
	} else if newpassword1, newpassword2 := api.HashPassword(m.Data.GetS("password1"), rt.Runtime.Salt), api.HashPassword(m.Data.GetS("password2"), rt.Runtime.Salt); newpassword1 != newpassword2 {
		update.ErrorSock(socket, "password change", "password mismatch")
	} else {
		log.Copy().Source().Trace(`about to engage`)
		account.Password = newpassword2
		if err := rt.Runtime.Root.Accounts.UpdatePassword(account); err != nil {
			update.ErrorSock(socket, "password change", err.Error())
		}
		socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(account.Username))
	}
}
