package internal

import (
	"net/url"

	"taylz.io/http"
)

func CORS(origins []string) func(h http.Handler) http.Handler {
	lookup := make(map[string]struct{})
	for _, origin := range origins {
		lookup[origin] = struct{}{}
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); origin == "" {
				w.WriteHeader(http.StatusUnauthorized)
			} else if url, err := url.Parse(origin); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else if _, ok := lookup[url.Host]; !ok {
				w.WriteHeader(http.StatusPreconditionFailed)
			} else {
				w.Header().Add("Access-Control-Allow-Origin", origin)
				w.Header().Add("Vary", "Origin")
				h.ServeHTTP(w, r)
			}
		})
	}
}
