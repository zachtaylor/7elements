package apiws

import (
	"ztaylor.me/cast"
)

// ping pushes a ping message to all sockets
func ping(rt *Runtime) {
	rt.WS.Message("/ping", cast.JSON{
		"ping":   rt.Runtime.Ping.Get(),
		"online": rt.Runtime.Sessions.Count(),
	})
}
