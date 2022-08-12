package internal

import (
	"taylz.io/http/user"
	"taylz.io/http/websocket"
)

func OnUser(server Server) user.ObserverFunc {
	return func(name string, newUser, oldUser *user.T) {
		if oldUser == nil && newUser != nil {
			go OnUserAdd(server, name)
		} else if oldUser != nil && newUser == nil {
			go OnUserRemove(server, name, oldUser.Sockets())
		} else {
			go OnUserWeird(server, name, oldUser, newUser)
		}
	}
}

func OnUserAdd(server Server, name string) {
	server.Log().Add("Name", name).Trace("new")
	server.Ping()
}

var dataNilAccount = websocket.NewMessage("/myaccount", nil).ShouldMarshal()

func OnUserRemove(server Server, name string, sockets []*websocket.T) {
	log := server.Log().Add("Name", name)
	for _, socket := range sockets {
		socket.WriteText(dataNilAccount)
	}
	if server.MatchMaker().Get(name) != nil {
		log.Warn("cancel queue")
		server.MatchMaker().Cancel(name)
	}
	log.Trace("old")
	server.Accounts().Remove(name)
	server.Ping()
}

func OnUserWeird(server Server, name string, oldUser, newUser *user.T) {
	server.Log().With(map[string]any{
		"Name": name,
		"Old":  oldUser,
		"New":  newUser,
	}).Warn("weird")
}
