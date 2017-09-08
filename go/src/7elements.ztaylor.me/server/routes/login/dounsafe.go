package login

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
	"net/http"
	"time"
	"ztaylor.me/log"
)

func DoUnsafe(account *accounts.Account, w http.ResponseWriter, r *http.Request, message string) {
	account.LastLogin = time.Now()
	if err := accounts.Store(account); err != nil {
		log.Add("Error", err).Error("cannot cache account")
		return
	}
	session := sessionman.GrantSession(account.Username)

	session.WriteSessionId(w)
	serverutil.WriteRedirectHome(w, message)
}
