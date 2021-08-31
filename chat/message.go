package chat

import "time"

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

func (m *Message) Data() map[string]interface{} {
	return map[string]interface{}{
		"user": m.Username,
		"chan": m.Channel,
		"msg":  m.Message,
		"date": m.Time.Format("2006-02-01"),
		"time": m.Time.Format("15:04:05"),
	}
}
