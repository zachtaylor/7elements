package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/http/websocket"
)

func Disconnect(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, _ *websocket.Message) {
		rt.Logger.New().Add("Socket", socket).Info()
		go rt.Ping()
	})
}
