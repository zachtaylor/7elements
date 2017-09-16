package server

import (
	"7elements.ztaylor.me/options"
	"net/http"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

var fileserver = http.FileServer(http.Dir(options.String("server-path")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if host := options.String("redirect-host"); host != "" && r.RequestURI == "/" && host != r.Host {
		log.Add("Hostname", host).Add("RequestHost", r.Host).Warn("page: redirected hostname")
		http.Redirect(w, r, r.URL.Scheme+"://"+host, 307)
		return
	}

	if session, _ := sessions.ReadRequestCookie(r); session != nil {
		session.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
