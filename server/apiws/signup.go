package apiws

import (
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/http/websocket"
)

// TODO write error messages to socket

func Signup(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Runtime.Root.Logger.New().Tag("apiws/signup").Add("Socket", socket.String())
		if socket.Session != nil {
			log.Warn("session exists")
		} else if username := m.Data.GetS("username"); username == "" {
			log.Warn("username missing")
		} else if !internal.CheckUsername(username) {
			log.Warn("username banned")
		} else if email := m.Data.GetS("email"); false {
		} else if a, err := rt.Runtime.Root.Accounts.Get(username); a != nil {
			log.Add("Error", err).Warn("account exists")
		} else if password1, password2 := internal.HashPassword(m.Data.GetS("password1"), rt.Runtime.Salt), internal.HashPassword(m.Data.GetS("password2"), rt.Runtime.Salt); password1 != password2 {
			log.Add("Error", err).Error("passwords don't match")
		} else if account, session, err := internal.Signup(rt.Runtime.Root, rt.Runtime.Sessions, username, email, password1); err != nil {
			log.Add("Error", err).Error("signup failed")
		} else {
			go ping(rt)
			redirect(socket, "/")
			socket.Session = session
			socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(account))
			log.Info()
		}
	})
}
