package out

import "ztaylor.me/cast"

// func Error(writer Writer, source, message string) {
// 	writer.WriteJSON(Build("/game/error", cast.JSON{
// 		"source":  source,
// 		"message": message,
// 	}))
// }

func GameError(t Target, source, message string) {
	t.Send("/game/error", cast.JSON{
		"source":  source,
		"message": message,
	})
}
