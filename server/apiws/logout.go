package apiws

import (
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Logout(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if len(socket.Name()) < 1 {
			rt.Logger.Warn("anon logout")
			socket.Write(wsout.ErrorJSON("vii", "you must log in to log out"))
			return
		}
		if session := rt.Sessions.GetName(socket.Name()); session != nil {
			rt.Sessions.Remove(session.ID())
		} else {

		}
		socket.SetName("")
		socket.WriteSync(wsout.MyAccount(nil).EncodeToJSON())
	})
}
