package out

import "ztaylor.me/cast"

func Error(t Target, source, message string) {
	t.Send("/error", cast.JSON{
		"source":  source,
		"message": message,
	})
}
