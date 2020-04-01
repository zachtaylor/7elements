package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
)

type Runtime struct {
	Runtime *api.Runtime
	WS      websocket.Service
}
