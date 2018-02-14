package chat

import (
	"sync"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

var CHANNEL_BUFFER_SIZE = 10

type Channel interface {
	AddSocket(*http.Socket)
	RemoveSocket(*http.Socket)
	AddMessage(*Message)
}

type channel struct {
	name    string
	sockets map[string]*http.Socket
	sync.Mutex
	history []*Message
}

func (c *channel) MessageJson(msg *Message) js.Object {
	return js.Object{
		"uri": "/chat",
		"data": js.Object{
			"channel":  c.name,
			"username": msg.Username,
			"message":  msg.Message,
		},
	}
}

func (c *channel) AddSocket(s *http.Socket) {
	c.Lock()
	c.sockets[s.Name()] = s
	for i := len(c.history) - 1; i >= 0; i-- {
		if msg := c.history[i]; msg != nil {
			s.WriteJson(c.MessageJson(msg))
		}
	}
	c.Unlock()
}

func (c *channel) RemoveSocket(s *http.Socket) {
	c.Lock()
	delete(c.sockets, s.Name())
	c.Unlock()
}

func (c *channel) AddMessage(msg *Message) {
	c.Lock()
	for i := len(c.history) - 1; i > 0; i-- {
		c.history[i] = c.history[i-1]
	}
	c.history[0] = msg
	json := c.MessageJson(msg)
	for _, s := range c.sockets {
		s.WriteJson(json)
	}
	c.Unlock()
}
