package runtime

import (
	"github.com/zachtaylor/7elements/db/accounts"
	"taylz.io/http/websocket"
)

func (t *T) Ping() {
	users, _ := accounts.Count(t.DB)
	bytes := websocket.NewMessage("/ping", websocket.MsgData{
		"ping":   len(t.Sockets.Keys()),
		"online": len(t.Sessions.Keys()),
		"users":  users,
	}).EncodeToJSON()
	for _, key := range t.Sockets.Keys() {
		if socket := t.Sockets.Get(key); socket != nil {
			socket.Write(bytes)
		}
	}
}
