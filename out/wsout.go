package out

import "taylz.io/http/websocket"

// import "taylz.io/http/websocket"

// func Ping(data websocket.JSON) []byte { return encode("/ping", data) }

// func Chat(data websocket.JSON) []byte { return encode("/chat", data) }

func Error(source, message string) []byte {
	return encode("/error", websocket.JSON{
		"source": source,
		"error":  message,
	})
}

func MyAccount(data websocket.JSON) []byte { return encode("/myaccount", data) }

// func Queue(data websocket.JSON) []byte { return encode("/queue", data) }

// // func MyAccountGame(id string) []byte {
// // 	return websocket.NewMessage("/myaccount/game", websocket.JSON{
// // 		"game": id,
// // 	})
// // }

// // Redirect sends a "/redirect" message
// //
// // path is expected to be like "/login" or something
// func Redirect(location string) []byte {
// 	return encode("/redirect", websocket.JSON{
// 		"location": location,
// 	})
// }
