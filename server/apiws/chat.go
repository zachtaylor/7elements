package apiws

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Chat(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if user, _, _ := rt.Users.GetSession(socket.SessionID()); user == nil {
			rt.Logger.Warn("anon chat")
			socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to chat"))
		} else if message, _ := m.Data["message"].(string); message == "" {
			rt.Logger.Add("User", user.Name()).Warn("message missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "message missing"))
		} else if channel, _ := m.Data["channel"].(string); channel == "" {
			rt.Logger.Add("User", user.Name()).Warn("channel missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "channel missing"))
		} else if err := rt.Chats.TryMessage(channel, user.Name(), message); err != nil {
			rt.Logger.Add("User", user.Name()).Error(err.Error())
			socket.WriteSync(wsout.ErrorJSON("vii", err.Error()))
		}
	})
}
