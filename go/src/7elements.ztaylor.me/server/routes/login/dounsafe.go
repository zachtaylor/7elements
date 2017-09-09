package login

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/options"
	"7elements.ztaylor.me/server/util"
	"net/http"
	"time"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func DoUnsafe(account *accounts.Account, w http.ResponseWriter, r *http.Request, message string) {
	account.LastLogin = time.Now()
	if err := accounts.Store(account); err != nil {
		log.Add("Error", err).Error("cannot cache account")
		return
	}
	session := sessions.Grant(account.Username, time.Duration(options.Int("session-life"))*time.Minute)
	account.SessionId = session.Id

	session.WriteCookie(w)
	serverutil.WriteRedirectHome(w, message)
}
