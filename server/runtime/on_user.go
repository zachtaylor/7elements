package runtime

import (
	"taylz.io/http/user"
	"taylz.io/http/websocket"
)

func (t *T) OnUser(name string, oldUser, newUser *user.T) {
	go t.onUser(name, oldUser, newUser)
}

func (t *T) onUser(name string, oldUser, newUser *user.T) {
	t.Log().With(websocket.MsgData{
		"Name":    name,
		"OldUser": oldUser != nil,
		"NewUser": newUser != nil,
	}).Trace("observe")
	if newUser == nil {
		t.Accounts.Remove(name)
	}
	go t.Ping()
}
