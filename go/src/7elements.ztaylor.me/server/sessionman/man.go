package sessionman

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/event"
	"errors"
	"net/http"
	"strconv"
	"time"
	"ztaylor.me/log"
)

var Cache = make(map[uint]*Session)

func GetSession(username string) *Session {
	if account := accounts.Test(username); account != nil {
		return Cache[account.SessionId]
	}
	return nil
}

func GrantSession(username string) *Session {
	if account := accounts.Test(username); account != nil {
		if account.SessionId > 0 {
			log.Add("OldSessionId", account.SessionId).Warn("session: overwrite account.sessionid")
			delete(Cache, account.SessionId)
		}

		session := NewSession()
		session.Username = username
		account.SessionId = session.Id
		Cache[session.Id] = session
		event.Fire("SessionGrant", session)
		log.Add("Username", username).Add("SessionId", session.Id).Debug("session: created")
		return session
	}

	log.Add("Username", username).Error("sessionman: account grantee must be cached")
	return nil
}

func RevokeSession(username string) {
	log := log.Add("Username", username)
	account := accounts.Test(username)
	if account == nil {
		log.Error("sessionman.RevokeSession: account missing")
		return
	} else if account.SessionId < 1 {
		log.Warn("sessionman.RevokeSession: sessionid missing")
		return
	}

	log.Add("SessionId", account.SessionId)
	session := Cache[account.SessionId]
	if session == nil {
		log.Warn("sessionman.RevokeSession: invalid sessionid")
		return
	}

	revoke(session)

	log.Add("LifeLeft", time.Now().Sub(session.Expire)).Debug("sessionman: revoke")
	event.Fire("SessionRevoke", session.Username)
}

func revoke(session *Session) {
	close(session.Done)
	delete(Cache, session.Id)
	if account := accounts.Test(session.Username); account != nil {
		account.SessionId = 0
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
				go revoke(session)
			}
		}
	}
}
