package apiws

import (
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func Chat(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Runtime.Root.Logger.New().Add("Session", socket.Session)
		if message := m.Data.GetS("message"); message == "" {
			log.Source().Warn("message missing")
		} else if message = cast.EscapeString(message); cast.OutCharset(message, SpeechCharset) {
			socket.Message("/error", cast.JSON{
				"error": "bad message content",
			})
		} else if channel := m.Data.GetS("channel"); channel == "" {
			log.Source().Warn("channel missing")
		} else if room := rt.Runtime.Chat.Get(channel); room == nil {
			log.Source().Warn("chat missing")
		} else if user := room.User(socket.ID); user == nil {
			log.Source().Warn("user missing")
		} else {
			msg := chat.NewMessage(user.Name, message)
			log.Source().Add("Message", msg).Info()
			room.AddMessage(msg)
		}
	})
}
