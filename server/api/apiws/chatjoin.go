package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	"ztaylor.me/http/json"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func ChatJoin(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		channel := m.Data.GetS("channel")
		log := rt.Root.Logger.New().With(log.Fields{
			"User":    m.User,
			"Channel": channel,
		}).Tag("apiws/chatjoin")

		if channel == "" {
			log.Warn("channel missing")
			return
		}

		if socket.Session == nil && rt.Chat.Get(channel) == nil {
			log.Debug("deny anon create")
			return
		}

		if chrm := rt.Chat.Get(channel); chrm == nil {
			log.Info("create")
		}

		chrm := rt.Chat.New(channel, 42)

		if chrm.Users[m.User] != nil {
			log.Warn("exists")
			return
		}

		user := newChatUser(channel, socket)
		chrm.User(user)

		history := make([]cast.JSON, 0)
		for _, msg := range chrm.History() {
			if msg != nil {
				history = append(history, newChatJSON(channel, msg))
			}
		}

		socket.Write(json.Encode(cast.JSON{
			"uri": "/chat/join",
			"data": cast.JSON{
				"channel":  channel,
				"username": m.User,
				"messages": history,
			},
		}))

		log.Info()
		select {
		case <-socket.Done():
			log.Debug("socket closed")
			delete(chrm.Users, m.User)
		case <-socket.Session.Done():
			log.Debug("session expired")
		}
	})
}
