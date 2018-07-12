package api

import (
	"elemen7s.com"
	"elemen7s.com/server/util"
	"net/http"
	"time"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

func GrantSession(a *vii.Account, w http.ResponseWriter, r *http.Request, message string) {
	a.LastLogin = time.Now()
	if err := vii.AccountService.UpdateLogin(a); err != nil {
		log.Add("Error", err).Error("login failed")
		return
	}
	session := zhttp.GrantSession(a.Username)
	a.SessionId = session.Id
	session.WriteCookie(w)
	serverutil.WriteRedirectHome(w, message)
}
