package apiws

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Email(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		email(rt, socket, m)
	})
}

func email(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
	log := rt.Logger.Add("Socket", socket.ID())

	if len(socket.SessionID()) < 1 {
		log.Warn("no session")
		log.Warn("anon update email")
		socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to change email"))
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
		log.Error("account missing")
		socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
		return
	}

	newemail, _ := m.Data["email"].(string)
	if newemail == "" {
		log.Add("Data", m.Data).Warn("email missing")
		socket.WriteSync(wsout.ErrorJSON("vii", "no new email"))
		return
	} else if err := api.CheckEmail(newemail); err != nil {
		log.Add("Error", err.Error()).Warn("bad email address")
		socket.WriteSync(wsout.ErrorJSON("vii", "bad email address"))
		return
	}
	log.Add("Email", newemail).Info()

	account.Email = newemail
	if err := accounts.UpdateEmail(rt.DB, account); err != nil {
		log.Add("Error", err.Error()).Error("update error")
		socket.WriteSync(wsout.ErrorJSON("vii", "bad email address"))
		return
	}

	socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
}
