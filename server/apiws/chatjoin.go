package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/http/websocket"
)

var SpeechCharset = charset.AlphaCapitalNumeric + " .-_+=!@$^&*()☺☻♥♦♣♠♂♀♪♫"

func ChatJoin(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		session := socket.Session
		channel := cast.EscapeString(m.Data.GetS("channel"))
		if cast.OutCharset(channel, SpeechCharset) {
			socket.Send("/error", cast.JSON{
				"error": "bad channel name",
			})
			return
		}
		log := rt.Logger.New().With(cast.JSON{
			"Session": session,
			"Channel": channel,
		})
		if channel == "" {
			log.Warn("channel missing")
			return
		}
		room := rt.Chats.Get(channel)
		if room == nil {
			log.Debug("room not found")
			if session == nil {
				socket.Send("/error", cast.JSON{"error": "cannot create room"})
				log.Warn("anon user new room")
				return
			}
			log.Info("create")
			room = rt.Chats.New(channel, 42)
		}
		user := room.AddUser(socket)
		history := make([]cast.JSON, 0)
		for _, msg := range room.History() {
			if msg != nil {
				history = append(history, newChatJSON(channel, msg))
			}
		}
		socket.Send("/chat/join", cast.JSON{
			// "userid":   socket.ID,
			"username": user.Name,
			"channel":  channel,
			"messages": history,
		})
		log.Info()
		select {
		case <-socket.DoneChan():
			log.Debug("socket closed")
		case <-socket.Session.Done():
			log.Debug("session expired")
		}
		room.Unsubscribe(user.Name)
	})
}
