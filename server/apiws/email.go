package apiws

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func Email(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		email(server, socket, m)
	})
}

func email(server internal.Server, socket *websocket.T, m *websocket.Message) {
	log := server.Log().Add("Socket", socket.ID())

	user := server.Users().GetWebsocket(socket)
	if user == nil {
		log.Warn("no session")
		socket.Write(websocket.MessageText, out.Error("vii", "no user"))
		return
	}
	log = log.Add("Session", user.Session())

	account := server.Accounts().Get(user.Session().Name())
	if account == nil {
		log.Error("account missing")
		socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
		return
	}

	newemail, _ := m.Data["email"].(string)
	if newemail == "" {
		log.Add("Data", m.Data).Warn("email missing")
		socket.Write(websocket.MessageText, out.Error("vii", "no new email"))
		return
	} else if err := api.CheckEmail(newemail); err != nil {
		log.Add("Error", err.Error()).Warn("bad email address")
		socket.Write(websocket.MessageText, out.Error("vii", "bad email address"))
		return
	}
	log.Add("Email", newemail).Info()

	account.Email = newemail
	if err := accounts.UpdateEmail(server.DB(), account); err != nil {
		log.Add("Error", err.Error()).Error("update error")
		socket.Write(websocket.MessageText, out.Error("vii", "bad email address"))
		return
	}

	socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
	socket.Write(websocket.MessageText, out.MyAccount(account.Data()))
}
