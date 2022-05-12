package apiws

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Email(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		email(server, socket, m)
	})
}

func email(server internal.Server, socket *websocket.T, m *websocket.Message) {
	log := server.Log().Add("Socket", socket.ID())

	if len(socket.SessionID()) < 1 {
		log.Warn("no session")
		log.Warn("anon update email")
		socket.WriteSync(wsout.Error("vii", "you must log in to change email"))
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
		log.Error("account missing")
		socket.WriteSync(wsout.Error("vii", "internal error"))
		return
	}

	newemail, _ := m.Data["email"].(string)
	if newemail == "" {
		log.Add("Data", m.Data).Warn("email missing")
		socket.WriteSync(wsout.Error("vii", "no new email"))
		return
	} else if err := api.CheckEmail(newemail); err != nil {
		log.Add("Error", err.Error()).Warn("bad email address")
		socket.WriteSync(wsout.Error("vii", "bad email address"))
		return
	}
	log.Add("Email", newemail).Info()

	account.Email = newemail
	if err := accounts.UpdateEmail(server.GetDB(), account); err != nil {
		log.Add("Error", err.Error()).Error("update error")
		socket.WriteSync(wsout.Error("vii", "bad email address"))
		return
	}

	socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
}
