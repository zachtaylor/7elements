package server

import (
	"net/http"
	"strings"
	"ztaylor.me/env"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

var fileserver = http.FileServer(http.Dir("www"))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if host := env.Default("HOST", ""); host != "" && r.RequestURI == "/" && host != r.Host {
		log.Add("Addr", r.RemoteAddr).Add("Hostname", host).Add("RequestHost", r.Host).Warn("page: redirected hostname")
		http.Redirect(w, r, "https://"+host, 307)
		return
	}

	remoteip := r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
	zhttp.Track(remoteip)

	if session, _ := zhttp.ReadRequestCookie(r); session != nil {
		session.Refresh()
		zhttp.TrackPair(session.Username, remoteip)
	}

	fileserver.ServeHTTP(w, r)
})
