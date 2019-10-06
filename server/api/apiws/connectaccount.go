package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func connectAccount(rt *api.Runtime, log *log.Entry, socket *websocket.T) {
	if session := socket.Session; session == nil {
		log.Debug("no session")
	} else {
		log.Add("Name", session.Name()).Debug("account data")
		pushJSON(socket, "/data/myaccount", rt.Root.AccountJSON(session.Name()))
	}
}
