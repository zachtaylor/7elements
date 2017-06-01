package jsbuilder

import (
	"net/http"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if content == "" {
		CreateContent()
	}

	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(content))
})
