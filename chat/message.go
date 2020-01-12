package chat

import (
	"time"

	"ztaylor.me/cast"
)

type Message struct {
	Username string
	Channel  string
	Message  string
	Time     time.Time
}

func NewMessage(username, message string) *Message {
	return &Message{
		Username: username,
		Message:  message,
		Time:     time.Now(),
	}
}

func (m *Message) JSON() cast.JSON {
	return cast.JSON{
		// "userid":    m.SocketID(),
		"username": m.Username,
		"message":  cast.EscapeString(m.Message),
		"time":     m.Time.Format("01-02 15:04:05"),
	}
}
