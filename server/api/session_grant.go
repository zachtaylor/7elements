package api

import (
	"net/http"
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/server/util"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

func GrantSession(a *vii.Account, w http.ResponseWriter, r *http.Request, message string) {
	a.LastLogin = time.Now()
	if err := vii.AccountService.UpdateLogin(a); err != nil {
		log.Add("Error", err).Error("api/session_grant: failed")
		return
	}
	session := zhttp.SessionService.Grant(a.Username)
	a.SessionId = session.ID
	session.WriteCookie(w)
	serverutil.WriteRedirectHome(w, message)
}