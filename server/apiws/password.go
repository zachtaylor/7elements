package apiws

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func Password(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		password(server, socket, m)
	})
}

func password(server internal.Server, socket *websocket.T, m *websocket.Message) {
	log := server.Log().Add("Socket", socket.ID())

	user := server.Users().GetWebsocket(socket)
	if user == nil {
		log.Warn("no session")
		socket.Write(websocket.MessageText, out.Error("vii", "you must log in to log out"))
		return
	}
	log = log.Add("Session", user.Session())

	account := server.Accounts().Get(user.Session().Name())
	if account == nil {
		log.Error("no account")
		socket.Write(websocket.MessageText, out.Error("vii", "no account"))
		return
	}
	log.Info()

	passbuff1, _ := m.Data["password1"].(string)
	if passbuff1 == "" {
		log.Warn("password1 missing")
		socket.Write(websocket.MessageText, out.Error("vii", "password1 missing"))
		return
	}
	pass1 := server.HashPassword(passbuff1)

	passbuff2, _ := m.Data["password2"].(string)
	if passbuff2 == "" {
		log.Warn("password2 missing")
		socket.Write(websocket.MessageText, out.Error("vii", "password2 missing"))
		return
	}
	pass2 := server.HashPassword(passbuff2)

	if pass1 != pass2 {
		log.Warn("password mismatch")
		socket.Write(websocket.MessageText, out.Error("vii", "password2 missing"))
		return
	}

	if err := accounts.UpdatePassword(server.DB(), account); err != nil {
		log.Add("Error", err.Error()).Error("update password")
		socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
		return
	}
	account.Password = pass2

	log.Info()
	socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
}
