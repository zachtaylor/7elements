package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/http/websocket"
)

func Email(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		email(rt, socket, m)
	})
}

func email(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
	log := rt.Logger.New().Add("Socket", socket).Add("Message", m)
	log.Info()
	if socket.Session == nil {
		out.Error(socket, "email change", "no session")
	} else if account, _ := rt.Accounts.Get(socket.Session.Name()); account == nil {
		out.Error(socket, "email change", "no account")
	} else if newemail := m.Data.GetS("email"); newemail == "" {
		out.Error(socket, "email change", "no new email")
	} else if cast.OutCharset(newemail, charset.AlphaCapitalNumeric+`-+@.`) {
		out.Error(socket, "email change", "bad email")
	} else {
		log.Copy().Trace(`about to engage`)
		account.Email = newemail
		if err := rt.Accounts.UpdateEmail(account); err != nil {
			out.Error(socket, "email change", err.Error())
		}
		socket.Send("/myaccount", rt.AccountJSON(account))
	}
}
