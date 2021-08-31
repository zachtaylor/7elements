package wsout

import "taylz.io/http/websocket"

func Chat(data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/chat", data)
}

func Error(source, message string) *websocket.Message {
	return websocket.NewMessage("/error", websocket.MsgData{
		"source": source,
		"error":  message,
	})
}

func ErrorJSON(source, message string) []byte { return Error(source, message).EncodeToJSON() }

func MyAccount(data websocket.MsgData) *websocket.Message {
	return websocket.NewMessage("/myaccount", data)
}

func MyAccountGame(id string) *websocket.Message {
	return websocket.NewMessage("/myaccount/game", websocket.MsgData{
		"game": id,
	})
}

// Redirect sends a "/redirect" message
//
// path is expected to be like "/login" or something
func Redirect(location string) *websocket.Message {
	return websocket.NewMessage("/redirect", websocket.MsgData{
		"location": location,
	})
}
