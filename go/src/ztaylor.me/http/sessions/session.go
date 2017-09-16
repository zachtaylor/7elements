package sessions

import (
	"github.com/cznic/mathutil"
	"net/http"
	"strconv"
	"sync"
	"time"
	"ztaylor.me/events"
)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999999, true)

type Session struct {
	Id       uint
	Username string
	Expire   time.Time
	Done     chan error
	sync.Mutex
	lifetime time.Duration
}

func New(lifetime time.Duration) *Session {
	return &Session{
		Id:       uint(sessionIdGen.Next()),
		Expire:   time.Now().Add(lifetime),
		lifetime: lifetime,
		Done:     make(chan error),
	}
}

func (session *Session) Refresh() {
	session.Expire = time.Now().Add(session.lifetime)
}

func (session *Session) Revoke() {
	close(session.Done)
	delete(Cache, session.Id)
	events.Fire("SessionRevoke", session.Username)
}

func (session *Session) WriteCookie(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(session.Id))+"; Path=/;")
}
