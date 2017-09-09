package server

import (
	"7elements.ztaylor.me/options"
	"net/http"
	"ztaylor.me/http/sessions"
)

var fileserver = http.FileServer(http.Dir(options.String("server-path")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if session, _ := sessions.ReadRequestCookie(r); session != nil {
		session.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
