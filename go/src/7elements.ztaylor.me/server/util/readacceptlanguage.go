package serverutil

import (
	"7elements.ztaylor.me/log"
	"net/http"
)

func ReadAcceptLanguage(r *http.Request) string {
	acceptLanguage := r.Header.Get("Accept-Language")
	acceptLanguage = acceptLanguage[0:5]
	if acceptLanguage == "" {
		acceptLanguage = "en-US"
		log.Warn("requests: accept-language header not identified, default en-US...")
	}

	return acceptLanguage
}
