package runtime

// import (
// 	"ztaylor.me/http/session"
// 	"ztaylor.me/http/websocket"
// )

// // T exports server data
// type T struct {
// 	Settings Settings
// 	Sessions *session.Cache
// 	Sockets  *websocket.Cache
// 	Players  *PlayerCache
// }

// // New creates a new server runtime
// func New(settings Settings, sessions *session.Cache, sockets *websocket.Cache) *T {
// 	return &T{
// 		Settings: settings,
// 		Sessions: sessions,
// 		Sockets:  sockets,
// 		Players:  NewPlayerCache(),
// 	}
// }

// func (t *T) Conn(socket *websocket.T) {
// 	if socket.Session == nil {
// 		t.Log().Add("Socket", socket).Trace("connect no session")
// 	} else {
// 		t.Log().Add("Socket", socket).Trace("connect session")
// 	}
// }

// // GetPlayer returns a player connected by the given key
// func (t *T) GetPlayer(name string) *Player {
// 	return t.Players.Get(name)
// }

// // func (t *T) sendPathMyAccount(socket *websocket.T, json cast.JSON) {
// // 	socket.Send("/myaccount", json)
// // }

// // func (t *T) waitSocketSession(socket *websocket.T) {
// // 	for socketDone, sessionDone := socket.DoneChan(), socket.Session.Done(); ; {
// // 		log := t.Log().Add("Socket", socket)
// // 		select {
// // 		case <-socketDone:
// // 			log.Debug("done")
// // 			return
// // 		case <-sessionDone:
// // 			log.Warn("session")
// // 			t.sendPathMyAccount(socket, nil)
// // 			socket.Session = nil
// // 			return
// // 		}
// // 	}
// // }
