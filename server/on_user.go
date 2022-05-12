package server

import (
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/user"
)

func OnUser(server internal.Server) user.CacheObserver {
	return func(name string, oldUser, newUser *user.T) {
		if oldUser == nil && newUser != nil {
			go OnUserAdd(server, name)
		} else if oldUser != nil && newUser == nil {
			go OnUserRemove(server, name, oldUser.Sockets())
		} else {
			go OnUserWeird(server, name, oldUser, newUser)
		}
	}
}

func OnUserAdd(server internal.Server, name string) {
	server.Log().Add("Name", name).Trace("new")
	server.Ping()
}

var dataNilAccount = wsout.MyAccount(nil).EncodeToJSON()

func OnUserRemove(server internal.Server, name string, sockets []string) {
	log := server.Log().Add("Name", name)
	for _, socketID := range sockets {
		if socket := server.GetWebsocketManager().Get(socketID); socket != nil {
			socket.WriteSync(dataNilAccount)
		}
	}
	if server.GetMatchMaker().Get(name) != nil {
		log.Warn("cancel queue")
		server.GetMatchMaker().Cancel(name)
	}
	log.Trace("old")
	server.GetAccounts().Remove(name)
	server.Ping()
}

func OnUserWeird(server internal.Server, name string, oldUser, newUser *user.T) {
	server.Log().With(map[string]any{
		"Name": name,
		"Old":  oldUser,
		"New":  newUser,
	}).Warn("weird")
}
