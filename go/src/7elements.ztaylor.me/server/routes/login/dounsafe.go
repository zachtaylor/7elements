package login

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
	"net/http"
	"time"
)

func DoUnsafe(account *SE.Account, w http.ResponseWriter, r *http.Request, message string) {
	account.LastLogin = time.Now()
	SE.Accounts.Cache[account.Username] = account
	session := sessionman.GrantSession(account.Username)

	session.WriteSessionId(w)
	log.Add("SessionId", session.Id)
	serverutil.WriteRedirectHome(w, message)
}