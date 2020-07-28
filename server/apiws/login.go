package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/http/websocket"
)

func Login(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		login(rt, socket, m)
	})
}

func login(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
	log := rt.Log().Add("Socket", socket)
	if socket.Session != nil {
		socket.Send("/error", cast.JSON{
			"error": "you are logged in!",
		})
		log.Warn("session exists")
	} else if username := m.Data.GetS("username"); username == "" {
		socket.Send("/error", cast.JSON{
			"error": "username missing",
		})
		log.Warn("username missing")
	} else if !cast.InCharset(username, charset.AlphaCapitalNumeric) {
		socket.Send("/error", cast.JSON{
			"error": "username not allowed: " + username,
		})
		log.Add("Name", username).Warn("username banned")
	} else if account, err := rt.Accounts.Get(username); account == nil {
		socket.Send("/error", cast.JSON{
			"error": "account missing",
		})
		log.Add("Error", err).Error("account missing")
	} else if password := internal.HashPassword(m.Data.GetS("password"), rt.PassSalt); password != account.Password {
		socket.Send("/error", cast.JSON{
			"error": "wrong password",
		})
		log.Add("Error", err).Error("wrong password")
	} else if p, err := rt.LoginSocket(socket, account); p == nil {
		socket.Send("/error", cast.JSON{
			"error": "(500) login failed", // censored
		})
		log.Add("Error", err).Error("login")
	} else { // accepted login
		log.Add("Name", account.Username).Info()
		socket.Session = p.Session
		redirect(socket, "/")
		connect(rt, socket)
		go rt.Ping()
	}
}

// WebsocketReceiver adapts websocket.T to cast.JSONWriter, to put in game.Seat.Receiver
type WebsocketReceiver struct {
	*websocket.T
}

func (ws *WebsocketReceiver) WriteJSON(data cast.JSON) {
	ws.Write(cast.BytesS(data.String()))
}

func (ws *WebsocketReceiver) String() string {
	return ws.T.String()
}
