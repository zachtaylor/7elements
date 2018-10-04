package api

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/http/ws"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func WSChatJoin() ws.Handler {
	newChatUser := func(socket *ws.Socket, channel string) func(msg *vii.ChatMessage) {
		return func(msg *vii.ChatMessage) {
			data := msg.JSON()
			data["channel"] = channel
			socket.WriteJson(js.Object{
				"uri":  "/chat",
				"data": data,
			})
		}
	}

	return ws.HandlerFunc(func(socket *ws.Socket, m *ws.Message) {
		channel := m.Data.Sval("channel")

		if channel == "" {
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("chat.join: channel missing")
			return
		}

		if socket.Session == nil && vii.ChatService.Get(channel) == nil {
			animate.Error(socket, `Join "`+channel+`"`, `cannot create chat`)
			log.WithFields(log.Fields{
				"User": m.User,
				"Data": m.Data,
			}).Warn("chat.join: session missing, chat missing")
			return
		}

		chat := vii.ChatService.New(channel)

		if chat != nil {
			log.WithFields(log.Fields{
				"User":    m.User,
				"Channel": channel,
			}).Info("chat.join: create")
		} else if chat = vii.ChatService.Get(channel); chat.Users[m.User] != nil {
			log.WithFields(log.Fields{
				"User":    m.User,
				"Channel": channel,
			}).Warn("chat.join: overwrite")
		}

		chat.Subscribe(m.User, newChatUser(socket, channel))

		history := make([]js.Object, 0)
		chat.Lock()
		for _, msg := range chat.History {
			if msg != nil {
				history = append(history, msg.JSON())
			}
		}
		chat.Unlock()

		socket.WriteJson(js.Object{
			"uri": "/chat/join",
			"data": js.Object{
				"channel":  channel,
				"username": m.User,
				"messages": history,
			},
		})

		log.WithFields(log.Fields{
			"User":    m.User,
			"Channel": channel,
		}).Info("chat.join")
	})
}
