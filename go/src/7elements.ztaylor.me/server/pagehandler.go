package server

import (
	"7elements.ztaylor.me/options"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
)

var fileserver = http.FileServer(http.Dir(options.String("server-path")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if session, _ := sessionman.ReadRequestCookie(r); session != nil {
		session.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
