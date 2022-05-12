package apiws

// import (
// 	"github.com/zachtaylor/7elements/server/internal"
// 	"github.com/zachtaylor/7elements/wsout"
// 	"taylz.io/http/websocket"
// )

// func Chat(server internal.Server) websocket.Handler {
// 	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
// 		if user, _, _ := server.GetUserManager().GetSocket(socket); user == nil {
// 			server.Log().Warn("anon chat")
// 			socket.WriteSync(wsout.Error("vii", "you must log in to chat"))
// 		} else if message, _ := m.Data["message"].(string); message == "" {
// 			server.Log().Add("User", user.Name()).Warn("message missing")
// 			socket.WriteSync(wsout.Error("vii", "message missing"))
// 		} else if channel, _ := m.Data["channel"].(string); channel == "" {
// 			server.Log().Add("User", user.Name()).Warn("channel missing")
// 			socket.WriteSync(wsout.Error("vii", "channel missing"))
// 		} else if err := runtime.GetChats().TryMessage(channel, user.Name(), message); err != nil {
// 			server.Log().Add("User", user.Name()).Error(err.Error())
// 			socket.WriteSync(wsout.Error("vii", err.Error()))
// 		}
// 	})
// }
