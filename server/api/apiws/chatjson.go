package apiws

import (
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/cast"
)

func newChatJSON(channel string, msg *chat.Message) cast.JSON {
	return cast.JSON{
		"channel":  channel,
		"username": msg.Username,
		"message":  msg.Message,
	}
}
