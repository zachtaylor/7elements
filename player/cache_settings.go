package player

import (
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/http/session"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

type CacheSettings struct {
	Logger   log.Service
	Sessions session.Service
	Sockets  *websocket.Cache
	Accounts account.Storer
}

func NewCache(logger log.Service, sessions session.Service, sockets *websocket.Cache, accounts account.Storer) *Cache {
	return &Cache{
		Settings: CacheSettings{
			Accounts: accounts,
			Logger:   logger,
			Sessions: sessions,
			Sockets:  sockets,
		},
		cache: make(map[string]*T),
	}
}
