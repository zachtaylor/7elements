package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func GrantSession(w http.ResponseWriter, r *http.Request, sessions *sessions.Service, a *vii.Account, message string) {
	a.LastLogin = time.Now()
	if err := vii.AccountService.UpdateLogin(a); err != nil {
		log.Add("Error", err).Error("api/session_grant")
		return
	}
	session := sessions.NewGrant(a.Username)
	a.SessionID = session.ID()
	log.Add("Account", a.Username).Add("SessionID", a.SessionID).Info("api/session_grant")
	session.WriteCookie(w)
	go waitForgetSession(session)
	w.Write([]byte(fmt.Sprintf(redirectHomeTpl, message)))
}

func waitForgetSession(session *sessions.T) {
	for {
		if _, ok := <-session.Done(); !ok {
			break
		}
	}
	vii.AccountService.Forget(session.Name())
	vii.AccountCardService.Forget(session.Name())
	vii.AccountDeckService.Forget(session.Name())
}

const redirectHomeTpl = `<html>
	<head>
		<title>Redirect</title>
	</head>
	<body>
		<img src="/img/banner-black.64px.png">
		<h3>%s</h3>
		<span style="font-size:21px">
			Redirecting to game in <b>30 s</b><br/>
			Click anywhere or press any key to go now
		</span>
		<script type="text/javascript">
			window.setTimeout(function() {
				window.location.pathname="/";
			}, 30000);
			document.addEventListener("click", function(e) {
				window.location.pathname="/";
			});
			document.addEventListener("keydown", function(e) {
				window.location.pathname="/";
			});
		</script>
	</body>
</html>
`
