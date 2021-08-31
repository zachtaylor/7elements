package runtime

import "taylz.io/http/session"

func (t *T) OnSession(id string, oldSession, newSession *session.T) {
	log := t.Logger.Add("Session", id)

	if oldSession == nil && newSession != nil {
		log.Add("Name", newSession.Name()).Info("open")
	} else if oldSession != nil && newSession == nil {
		log.Add("Name", oldSession.Name()).Info("close")
	} else {
		log.Add("Old", oldSession).Add("New", newSession).Warn("weird")
	}
}
