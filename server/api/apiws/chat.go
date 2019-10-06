package apiws

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
)

func Chat(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Root.Logger.New().Tag("api/ws/chat").Add("Username", m.User)
		if message := m.Data.GetS("message"); message == "" {
			log.Add("Data", m.Data.String()).Warn("message missing")
		} else if channel := m.Data.GetS("channel"); channel == "" {
			log.Add("Data", m.Data.String()).Warn("channel missing")
		} else if chrm := rt.Chat.Get(channel); chrm == nil {
			log.Add("Data", m.Data.String()).Warn("chat missing")
		} else {
			msg := chat.NewMessage(m.User, message)
			go chrm.AddMessage(msg)
			log.Add("Channel", channel).Add("Message", msg.Message).Info("chat")
		}
	})
}
