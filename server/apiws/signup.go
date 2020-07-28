package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/http/websocket"
)

// TODO write error messages to socket

func Signup(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Log().Tag("apiws/signup").Add("Socket", socket.String())
		password1 := internal.HashPassword(m.Data.GetS("password1"), rt.PassSalt)
		password2 := internal.HashPassword(m.Data.GetS("password2"), rt.PassSalt)
		if socket.Session != nil {
			log.Warn("session exists")
		} else if username := m.Data.GetS("username"); username == "" {
			log.Warn("username missing")
		} else if !internal.CheckUsername(username) {
			log.Warn("username banned")
		} else if email := m.Data.GetS("email"); false {
		} else if a, err := rt.Accounts.Get(username); a != nil {
			log.Add("Error", err).Warn("account exists")
		} else if password1 != password2 {
			log.Add("Error", err).Error("passwords don't match")
		} else if player, err := rt.Players.Signup(username, email, password1); err != nil {
			log.Add("Error", err).Error("signup failed")
		} else {
			go rt.Ping()
			redirect(socket, "/")
			// socket.Session = player.Session
			socket.Send("/myaccount", rt.AccountJSON(player.Account))
			log.Info()
		}
	})
}
