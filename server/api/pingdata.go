package api

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"ztaylor.me/cast"
)

func PingData(rt *runtime.T) cast.JSON {
	return cast.JSON{
		"online": rt.Sessions.Count(),
	}
}
