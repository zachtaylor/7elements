package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

type Runtime struct {
	Runtime *api.Runtime
	WS      websocket.Service
}

func (rt *Runtime) SendPing() {
	rt.WS.Message("/ping", cast.JSON{
		"ping":   rt.Runtime.Ping.Get(),
		"online": rt.Runtime.Sessions.Count(),
	})
}
