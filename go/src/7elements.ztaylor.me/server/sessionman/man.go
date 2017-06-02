package sessionman

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/event"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/options"
	"net/http"
	"strconv"
	"time"
)

func GrantSession(username string) *Session {
	if account := SE.Accounts.Cache[username]; account != nil {
		if account.SessionId > 0 {
			log.Add("OldSessionId", account.SessionId).Warn("session: overwrite account.sessionid")
			delete(cache, account.SessionId)
		}

		session := &Session{
			Username: account.Username,
			Id:       uint(sessionIdGen.Next()),
			Expire:   time.Now().Add(time.Duration(options.Int("session-life")) * time.Second),
		}

		account.SessionId = session.Id
		cache[session.Id] = session
		event.Fire("GrantSession", session)
		log.Add("Username", username).Add("SessionId", session.Id).Debug("session: created")
		return session
	}

	log.Add("Username", username).Error("sessionman: account grantee must be cached")
	return nil
}

func RevokeSession(username string) {
	log.Add("Username", username)

	if account := SE.Accounts.Cache[username]; account != nil {
		if account.SessionId < 1 {
			log.Warn("sessionman: revoke: account does not have session")
			return
		}

		log.Add("SessionId", account.SessionId)

		session := cache[account.SessionId]
		if session == nil {
			log.Warn("sessionman: revoke: account points to invalid session")
			return
		}

		log.Add("LifeLeft", time.Now().Sub(session.Expire))

		delete(cache, account.SessionId)
		account.SessionId = 0
		event.Fire("RevokeSession", session.Username)
		log.Debug("sessionman: revoke")
	} else {
		log.Error("sessionman: revoke: account must be cached")
	}
}

func ReadRequestCookie(r *http.Request) (*Session, error) {
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			if session := cache[uint(sessionId)]; session != nil {
				return session, nil
			} else if sessionId == 0 {
				return nil, nil
			} else {
				return nil, Errors.InvalidSessionId
			}
		} else {
			return nil, Errors.CookieParse
		}
	} else {
		return nil, Errors.NoCookie
	}
}

func EraseSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId=0; Path=/;")
}
