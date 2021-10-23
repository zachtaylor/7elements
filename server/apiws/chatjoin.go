package apiws

// import (
// 	"strings"

// 	"github.com/zachtaylor/7elements/server/runtime"
// 	"github.com/zachtaylor/7elements/wsout"
// 	"taylz.io/http/websocket"
// 	"taylz.io/keygen/charset"
// )

// const SpeechCharset = charset.AlphaCapitalNumeric + " ,.-_+=!@$^&*()☺☻♥♦♣♠♂♀♪♫"

// func ChatJoin(rt *runtime.T) websocket.Handler {
// 	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
// 		if len(socket.Name()) < 1 {
// 			rt.Logger.Warn("anon chat")
// 			socket.Write(wsout.ErrorJSON("vii", "you must log in to chat"))
// 		} else if channel, _ := m.Data["channel"].(string); len(channel) < 1 {
// 			rt.Logger.Add("User", socket.Name()).Warn("channel missing")
// 			socket.Write(wsout.ErrorJSON("vii", "channel missing"))
// 		} else if symbols := strings.Trim(channel, charset.AlphaCapitalNumeric); len(symbols) > 1 {
// 			rt.Logger.Add("User", socket.Name()).Add("Symbols", symbols).Warn("bad channel name: ", channel)
// 			socket.Write(wsout.ErrorJSON("vii", "bad channel name"))
// 		} else if room := rt.Chats.Get(channel); room == nil {
// 			rt.Logger.Add("User", socket.Name()).Info("create")
// 			room = rt.Chats.New(channel, 42)
// 		} else {
// 			rt.Logger.With(map[string]interface{}{
// 				"User": socket.Name(),
// 				"Room": channel,
// 			}).Info("join")
// 			room.AddUser(socket.Name())

// 			// socket.Send("/chat/join", websocket.MsgData{
// 			// // "userid":   socket.ID,
// 			// 	"username": user.Session.Name,
// 			// 	"channel":  channel,
// 			// 	"messages": history,
// 			// })

// 			for _, data := range room.History() {
// 				socket.WriteSync(data)
// 			}
// 		}
// 	})
// }
