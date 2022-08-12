package apiws

// import (
// 	"github.com/zachtaylor/7elements/server/internal"
// 	"github.com/zachtaylor/7elements/wsout"
// 	"taylz.io/http/websocket"
// )

// func Chat(server internal.Server) websocket.MessageHandler {
// 	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
// 		if user, _, _ := server.Users().GetSocket(socket); user == nil {
// 			server.Log().Warn("anon chat")
// 			socket.Write(websocket.MessageText, out.Error("vii", "you must log in to chat"))
// 		} else if message, _ := m.Data["message"].(string); message == "" {
// 			server.Log().Add("User", user.Name()).Warn("message missing")
// 			socket.Write(websocket.MessageText, out.Error("vii", "message missing"))
// 		} else if channel, _ := m.Data["channel"].(string); channel == "" {
// 			server.Log().Add("User", user.Name()).Warn("channel missing")
// 			socket.Write(websocket.MessageText, out.Error("vii", "channel missing"))
// 		} else if err := runtime.GetChats().TryMessage(channel, user.Name(), message); err != nil {
// 			server.Log().Add("User", user.Name()).Error(err.Error())
// 			socket.Write(websocket.MessageText, out.Error("vii", err.Error()))
// 		}
// 	})
// }
