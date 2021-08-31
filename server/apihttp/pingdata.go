package apihttp

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http/websocket"
)

func PingData(rt *runtime.T) websocket.MsgData {
	return websocket.MsgData{
		"online": len(rt.Sessions.Keys()),
	}
}
