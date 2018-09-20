package server // import "github.com/zachtaylor/7elements/server"

import (
	"net/http"

	"ztaylor.me/http/mux"
)

func New(fs http.FileSystem, dbsalt string) *mux.Mux {
	mux := mux.NewMux()
	Routes(mux, fs, dbsalt)
	return mux
}
