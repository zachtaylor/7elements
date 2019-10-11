package chat

import (
	"strings"
	"time"

	"ztaylor.me/cast"
)

type Message struct {
	Username string
	Message  string
	Time     time.Time
}

func NewMessage(username string, message string) *Message {
	return &Message{
		Username: username,
		Message:  message,
		Time:     time.Now(),
	}
}

// func (m *Message) String() string {
// 	return m.JSON().String()
// }

func (m *Message) JSON() cast.JSON {
	return cast.JSON{
		"username": m.Username,
		"message":  strings.Replace(m.Message, "\"", "\\\"", -1),
		"time":     m.Time.Format("01-02 15:04:05"),
	}
}
