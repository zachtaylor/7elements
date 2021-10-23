package runtime

import "taylz.io/http/session"

func (rt *T) OnSession(id string, oldSession, newSession *session.T) {
	go rt.onSession(id, oldSession, newSession)
}

func (rt *T) onSession(id string, oldSession, newSession *session.T) {
	log := rt.Logger.Add("Session", id)
	if oldSession == nil && newSession != nil {
		log.Add("Name", newSession.Name()).Info("login ")
	} else if oldSession != nil && newSession == nil {
		log.Add("Name", oldSession.Name()).Info("logout")
	} else {
		log.Add("Old", oldSession).Add("New", newSession).Warn("weird")
	}
}
