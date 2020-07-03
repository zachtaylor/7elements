package player

import (
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/cast"
	"ztaylor.me/http/session"
)

// T represents a live player
type T struct {
	Settings Settings
	Session  *session.T
	Account  *account.T
	GameID   string
	conns    map[string]bool
	lock     cast.Mutex
}

func (player *T) Name() string {
	if player.Account == nil {
		return "anon"
	}
	return player.Account.Username
}

func (player *T) AddConn(key string) {
	player.lock.Lock()
	player.conns[key] = true
	player.lock.Unlock()
}

func (player *T) RemConn(key string) {
	player.lock.Lock()
	delete(player.conns, key)
	player.lock.Unlock()
}

func (player *T) Send(uri string, data cast.JSON) {
	player.lock.Lock()
	for id := range player.conns {
		if conn := player.Settings.Sockets.Get(id); conn == nil {

		} else {
			conn.Send(uri, data)
		}
	}
	player.lock.Unlock()
}
