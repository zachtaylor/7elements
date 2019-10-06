package api

import "ztaylor.me/cast"

func PingData(rt *Runtime) cast.JSON {
	return cast.JSON{
		"online": rt.Sessions.Count(),
	}
}
