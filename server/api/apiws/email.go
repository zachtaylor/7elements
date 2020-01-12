package apiws

import (
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/http/websocket"
)

func Email(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		email(rt, socket, m)
	})
}

func email(rt *Runtime, socket *websocket.T, m *websocket.Message) {
	log := rt.Runtime.Root.Logger.New().Add("Socket", socket).Add("Message", m)
	log.Copy().Source().Info()
	session := socket.Session
	if session == nil {
		update.ErrorSock(socket, "email change", "no session")
	} else if account := rt.Runtime.Root.Accounts.Test(session.Name()); account == nil {
		update.ErrorSock(socket, "email change", "no account")
	} else if newemail := m.Data.GetS("email"); newemail == "" {
		update.ErrorSock(socket, "email change", "no new email")
	} else if cast.OutCharset(newemail, charset.AlphaCapitalNumeric+`-+@.`) {
		update.ErrorSock(socket, "email change", "bad email")
	} else {
		log.Copy().Source().Trace(`about to engage`)
		account.Email = newemail
		if err := rt.Runtime.Root.Accounts.UpdateEmail(account); err != nil {
			update.ErrorSock(socket, "email change", err.Error())
		}
		socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(account.Username))
	}
}
