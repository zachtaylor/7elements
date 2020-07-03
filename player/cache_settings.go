package player

import (
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/http/session"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

type CacheSettings struct {
	Accounts account.Service
	Logger   log.Service
	Sessions session.Service
	Sockets  websocket.Service
}

func NewCacheSettings(accounts account.Service, logger log.Service, sessions session.Service, sockets websocket.Service) CacheSettings {
	return CacheSettings{
		Accounts: accounts,
		Logger:   logger,
		Sessions: sessions,
		Sockets:  sockets,
	}
}
