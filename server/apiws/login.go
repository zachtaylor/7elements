package apiws

import (
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/http/websocket"
)

func Login(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		login(rt, socket, m)
	})
}

func login(rt *Runtime, socket *websocket.T, m *websocket.Message) {
	log := rt.Runtime.Root.Logger.New().Add("Socket", socket)
	if socket.Session != nil {
		socket.Message("/error", cast.JSON{
			"error": "you are logged in!",
		})
		log.Source().Warn("session exists")
	} else if username := m.Data.GetS("username"); username == "" {
		socket.Message("/error", cast.JSON{
			"error": "username missing",
		})
		log.Warn("username missing")
	} else if !cast.InCharset(username, charset.AlphaCapitalNumeric) {
		socket.Message("/error", cast.JSON{
			"error": "username not allowed: " + username,
		})
		log.Add("Name", username).Warn("username banned")
	} else if account, err := rt.Runtime.Root.Accounts.Get(username); account == nil {
		socket.Message("/error", cast.JSON{
			"error": "account missing",
		})
		log.Add("Error", err).Error("account missing")
	} else if password := internal.HashPassword(m.Data.GetS("password"), rt.Runtime.Salt); password != account.Password {
		socket.Message("/error", cast.JSON{
			"error": "wrong password",
		})
		log.Add("Error", err).Error("wrong password")
	} else if s, err := internal.Login(rt.Runtime.Root, rt.Runtime.Sessions, account); s == nil {
		socket.Message("/error", cast.JSON{
			"error": "(500) login failed", // censored
		})
		log.Add("Error", err).Error("login")
	} else { // accepted login
		log.Add("Name", account.Username).Source().Info()
		rt.Runtime.Root.Accounts.Cache(account)
		socket.Session = s
		redirect(socket, "/")
		connect(rt, socket)
	}
}

// WebsocketReceiver adapts websocket.T to cast.JSONWriter, to put in game.Seat.Receiver
type WebsocketReceiver struct {
	*websocket.T
}

func (ws *WebsocketReceiver) WriteJSON(data cast.JSON) {
	ws.Send(cast.BytesS(data.String()))
}

func (ws *WebsocketReceiver) String() string {
	return ws.T.String()
}
