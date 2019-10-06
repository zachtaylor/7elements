package apiws

import (
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/http/websocket"
)

func newChatUser(channel string, socket *websocket.T) *chat.User {
	return &chat.User{
		Name: socket.Session.Name(),
		Send: func(msg *chat.Message) {
			pushJSON(socket, "/chat", newChatJSON(channel, msg))
		},
	}
}
