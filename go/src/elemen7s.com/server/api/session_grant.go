package api

import (
	"elemen7s.com/accounts"
	"elemen7s.com/server/util"
	"net/http"
	"time"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

func GrantSession(a *accounts.Account, w http.ResponseWriter, r *http.Request, message string) {
	a.LastLogin = time.Now()
	if err := accounts.Store(a); err != nil {
		log.Add("Error", err).Error("cannot cache account")
		return
	}
	session := zhttp.GrantSession(a.Username)
	a.SessionId = session.Id
	session.WriteCookie(w)
	serverutil.WriteRedirectHome(w, message)
}
