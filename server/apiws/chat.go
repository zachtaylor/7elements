package apiws

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func Chat(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Logger.New().Add("Session", socket.Session)
		if message := m.Data.GetS("message"); message == "" {
			log.Warn("message missing")
		} else if message = cast.EscapeString(message); cast.OutCharset(message, SpeechCharset) {
			socket.Send("/error", cast.JSON{
				"error": "bad message content",
			})
		} else if channel := m.Data.GetS("channel"); channel == "" {
			log.Warn("channel missing")
		} else if room := rt.Chats.Get(channel); room == nil {
			log.Warn("chat missing")
		} else if user := room.User(socket.ID()); user == nil {
			log.Warn("user missing")
		} else {
			msg := chat.NewMessage(user.Name, message)
			log.Add("Message", msg).Info()
			room.AddMessage(msg)
		}
	})
}
