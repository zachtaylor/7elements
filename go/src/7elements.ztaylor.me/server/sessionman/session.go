package sessionman

import (
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/options"
	"github.com/cznic/mathutil"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Session struct {
	Id        uint
	Username  string
	Websocket *websocket.Conn
	Expire    time.Time
	sync.Mutex
}

var cache = make(map[uint]*Session)
var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

func (session *Session) Refresh() {
	session.Expire = time.Now().Add(time.Duration(options.Int("session-life")) * time.Second)
}

func (session *Session) Send(message string) {
	if session.Websocket != nil {
		websocket.Message.Send(session.Websocket, message)
	} else {
		log.Add("SessionId", session.Id).Add("Username", session.Username).Warn("session.Frame: websocket missing")
	}
}

func (session *Session) WriteSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(session.Id))+"; Path=/;")
}

func SessionClock() {
	for now := range time.Tick(1 * time.Second) {
		for _, session := range cache {
			if session.Expire.Before(now) {
				RevokeSession(session.Username)

				if session.Websocket != nil {
					session.Websocket.Close()
				}
			}
		}
	}
}
