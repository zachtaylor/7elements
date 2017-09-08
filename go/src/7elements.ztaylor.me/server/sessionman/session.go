package sessionman

import (
	"7elements.ztaylor.me/options"
	"github.com/cznic/mathutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999999, true)

type Session struct {
	Id       uint
	Username string
	Expire   time.Time
	Done     chan error
	sync.Mutex
}

func NewSession() *Session {
	return &Session{
		Id:     uint(sessionIdGen.Next()),
		Expire: time.Now().Add(time.Duration(options.Int("session-life")) * time.Second),
		Done:   make(chan error),
	}
}

func (session *Session) Refresh() {
	session.Expire = time.Now().Add(time.Duration(options.Int("session-life")) * time.Second)
}

func (session *Session) WriteSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(session.Id))+"; Path=/;")
}
