package api

import (
	"net/http"
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/server/util"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func GrantSession(w http.ResponseWriter, r *http.Request, a *vii.Account, message string) {
	a.LastLogin = time.Now()
	if err := vii.AccountService.UpdateLogin(a); err != nil {
		log.Add("Error", err).Error("api/session_grant")
		return
	}
	session := sessions.Service.Grant(a.Username)
	a.SessionID = session.ID
	log.Add("Account", a.Username).Add("SessionID", a.SessionID).Info("api/session_grant")
	sessions.WriteCookie(w, session)
	serverutil.WriteRedirectHome(w, message)
}
