package api

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/http/ws"
	"ztaylor.me/log"
)

func WSChat() ws.Handler {
	return ws.HandlerFunc(func(socket *ws.Socket, m *ws.Message) {
		if message := m.Data.Sval("message"); message == "" {
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("chat: message missing")
		} else if channel := m.Data.Sval("channel"); channel == "" {
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("chat: channel missing")
		} else if chat := vii.ChatService.Get(channel); chat == nil {
			log.WithFields(log.Fields{
				"User":    m.User,
				"Message": message,
				"Channel": channel,
			}).Warn("chat: missing")
		} else {
			msg := vii.NewChatMessage(m.User, message)
			go chat.AddMessage(msg)
			log.WithFields(log.Fields{
				"User":    m.User,
				"Message": message,
				"Channel": channel,
			}).Info("chat")
		}
	})
}
