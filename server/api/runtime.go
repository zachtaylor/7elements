package api

import (
	"net/http"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/http/session"
)

type Runtime struct {
	Root       *vii.Runtime
	Games      game.Service
	Chat       chat.Service
	Salt       string
	Sessions   session.Service
	FileSystem http.FileSystem
}

// func NewRuntime(root *vii.Runtime)
