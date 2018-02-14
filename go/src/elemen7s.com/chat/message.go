package chat

import (
	"time"
)

type Message struct {
	Username string
	Message  string
	time.Time
}

func NewMessage(username string, message string) *Message {
	return &Message{
		Username: username,
		Message:  message,
		Time:     time.Now(),
	}
}
