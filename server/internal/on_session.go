package internal

import "taylz.io/http/session"

// OnSession creates a closure for the internal.Server definition to react to a change in session cache
func OnSession(server Server) session.ObserverFunc {
	return func(id string, newSession, oldSession *session.T) {
		go OnSessionServer(server, id, newSession, oldSession) // surrender hotpath instantly
	}
}

func OnSessionServer(server Server, id string, newSession, oldSession *session.T) {
	if oldSession == nil && newSession != nil {
		server.Log().With(map[string]any{
			"ID":   id,
			"Name": newSession.Name(),
		}).Debug("login")
	} else if oldSession != nil && newSession == nil {
		server.Log().With(map[string]any{
			"ID":   id,
			"Name": oldSession.Name(),
		}).Debug("logout")
	} else {
		server.Log().With(map[string]any{
			"ID":  id,
			"New": newSession,
			"Old": oldSession,
		}).Warn("weird")
	}
}
