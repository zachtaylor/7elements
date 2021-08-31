package apiws

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Chat(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if len(socket.Name()) < 1 {
			rt.Logger.Warn("anon chat")
			socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to chat"))
		} else if message, _ := m.Data["message"].(string); message == "" {
			rt.Logger.Add("User", socket.Name()).Warn("message missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "message missing"))
		} else if channel, _ := m.Data["channel"].(string); channel == "" {
			rt.Logger.Add("User", socket.Name()).Warn("channel missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "channel missing"))
		} else if err := rt.Chats.TryMessage(channel, socket.Name(), message); err != nil {
			rt.Logger.Add("User", socket.Name()).Error(err.Error())
			socket.WriteSync(wsout.ErrorJSON("vii", err.Error()))
		}
	})
}
