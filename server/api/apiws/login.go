package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	"ztaylor.me/http/json"
	"ztaylor.me/http/websocket"
)

func Login(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Root.Logger.New().Tag("apiws/login").Add("Remote", socket.Key())
		if socket.Session != nil {
			pushJSON(socket, "/error", cast.JSON{
				"error": "you are logged in!",
			})
			log.Warn("session exists")
		} else if username := m.Data.GetS("username"); username == "" {
			pushJSON(socket, "/error", cast.JSON{
				"error": "username missing",
			})
			log.Warn("username missing")
		} else if !api.CheckUsername(username) {
			pushJSON(socket, "/error", cast.JSON{
				"error": "username not allowed: " + username,
			})
			log.Add("Name", username).Warn("username banned")
		} else if account, err := rt.Root.Accounts.Get(username); account == nil {
			pushJSON(socket, "/error", cast.JSON{
				"error": "account missing",
			})
			log.Add("Error", err).Error("account missing")
		} else if password := api.HashPassword(m.Data.GetS("password"), rt.Salt); password != account.Password {
			pushJSON(socket, "/error", cast.JSON{
				"error": "wrong password",
			})
			log.Add("Error", err).Error("wrong password")
		} else if s, err := api.Login(rt, account); s == nil {
			pushJSON(socket, "/error", cast.JSON{
				"error": "(500) login failed", // censored
			})
			log.Add("Error", err).Error("login")
		} else { // accepted login
			rt.Root.Accounts.Cache(account)
			socket.Session = s
			log.Add("Name", account.Username).Info("accept")
			pushJSON(socket, "/data/ping", api.PingData(rt))
			pushRedirectJSON(socket, "/")
			connectAccount(rt, log, socket)
			connectGame(rt, log, socket)
		}
	})
}

// WebsocketReceiver adapts websocket.T to cast.JSONWriter, to put in game.Seat.Receiver
type WebsocketReceiver struct {
	*websocket.T
}

func (ws *WebsocketReceiver) WriteJSON(data cast.JSON) {
	ws.Write(json.Encode(data))
}

func (ws *WebsocketReceiver) String() string {
	return ws.T.String()
}
