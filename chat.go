package vii

import (
	"sync"
	"time"

	"ztaylor.me/js"
)

type Chat struct {
	Key     string
	Users   map[string]func(*ChatMessage)
	History []*ChatMessage
	sync.Mutex
}

type ChatMessage struct {
	Username string
	Message  string
	Time     time.Time
}

func NewChat(key string, messageBuffer int) *Chat {
	return &Chat{
		Key:     key,
		Users:   make(map[string]func(*ChatMessage)),
		History: make([]*ChatMessage, messageBuffer),
	}
}

func NewChatMessage(username string, message string) *ChatMessage {
	return &ChatMessage{
		Username: username,
		Message:  message,
		Time:     time.Now(),
	}
}

// func (m *ChatMessage) String() string {
// 	return m.JSON().String()
// }

func (m *ChatMessage) JSON() js.Object {
	return js.Object{
		"username": m.Username,
		"message":  m.Message,
		"time":     m.Time.Format("01-02 15:04:05"),
	}
}

func (c *Chat) AddMessage(msg *ChatMessage) {
	c.Lock()
	for i := len(c.History) - 1; i > 0; i-- {
		c.History[i] = c.History[i-1]
	}
	c.History[0] = msg
	for _, user := range c.Users {
		user(msg)
	}
	c.Unlock()
}

func (c *Chat) Subscribe(name string, user func(*ChatMessage)) {
	c.Lock()
	c.Users[name] = user
	c.Unlock()
}

func (c *Chat) Unsubscribe(username string) {
	c.Lock()
	delete(c.Users, username)
	c.Unlock()
}

type MemChatService map[string]*Chat

func (service MemChatService) Get(key string) *Chat {
	return service[key]
}

func (service MemChatService) New(key string) *Chat {
	if service[key] != nil {
		return nil
	}
	chat := NewChat(key, 12)
	service[key] = chat
	return chat
}

var ChatService interface {
	Get(string) *Chat
	New(string) *Chat
} = MemChatService{}
