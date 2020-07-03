package chat

import "ztaylor.me/http/websocket"

type User struct {
	Name   string
	Socket *websocket.T
}

func NewUser(name string, socket *websocket.T) *User {
	return &User{
		Name:   name,
		Socket: socket,
	}
}
