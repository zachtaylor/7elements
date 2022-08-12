package game

import (
	"nhooyr.io/websocket"
	"taylz.io/http/user"
)

// Writer is a wrapper for sending data
type Writer interface {
	Name() string
	Done() <-chan struct{}
	Write([]byte) error
}

type UserWriter struct {
	User *user.T
}

func (w UserWriter) Name() string { return w.User.Session().Name() }

func (w UserWriter) Done() <-chan struct{} { return w.User.Done() }

func (w UserWriter) Write(buf []byte) error { return w.User.Write(websocket.MessageText, buf) }
