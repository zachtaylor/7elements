package server

import (
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/session"
)

// OnSession creates a closure for the internal.Server definition to react to a change in session cache
func OnSession(server internal.Server) session.CacheObserver {
	return func(id string, oldSession, newSession *session.T) {
		go OnSessionServer(server, id, oldSession, newSession) // surrender hotpath instantly
	}
}

func OnSessionServer(server internal.Server, id string, oldSession, newSession *session.T) {
	if oldSession == nil && newSession != nil {
		server.Log().Info("login", id, newSession.Name())
	} else if oldSession != nil && newSession == nil {
		server.Log().Info("logout", id, oldSession.Name())
	} else {
		server.Log().Info("logx", id, oldSession, newSession)
	}
}
