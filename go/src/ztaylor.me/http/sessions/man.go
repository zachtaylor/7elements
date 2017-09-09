package sessions

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"ztaylor.me/events"
	"ztaylor.me/log"
)

var Cache = make(map[uint]*Session)

func Get(username string) *Session {
	for _, session := range Cache {
		if username == session.Username {
			return session
		}
	}
	return nil
}

func Grant(username string, lifetime time.Duration) *Session {
	session := New(lifetime)
	session.Username = username
	Cache[session.Id] = session
	events.Fire("SessionGrant", session)
	log.Add("Username", username).Add("SessionId", session.Id).Add("Lifetime", lifetime).Debug("session: created")
	return session
}

func Revoke(username string) {
	log := log.Add("Username", username)

	if session := Get(username); session == nil {
		log.Warn("sessions.revoke: session missing")
	} else {
		log.Add("Username", session.Username).Add("Expiration", session.Expire.Unix()).Add("TimeNow", time.Now().Unix()).Info("session expired")
		session.Expire = time.Now()
	}
}

func ReadRequestCookie(r *http.Request) (*Session, error) {
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			if session := Cache[uint(sessionId)]; session != nil {
				return session, nil
			} else if sessionId == 0 {
				return nil, nil
			} else {
				return nil, errors.New("invalid cookie#" + sessionCookie.Value)
			}
		} else {
			return nil, errors.New("cookie format")
		}
	} else {
		return nil, errors.New("session missing")
	}
}

func EraseSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId=0; Path=/;")
}

func SessionClock() {
	for now := range time.Tick(1 * time.Second) {
		for _, session := range Cache {
			if session.Expire.Before(now) {
				go session.Revoke()
			}
		}
	}
}
