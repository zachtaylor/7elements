package animate // import "github.com/zachtaylor/7elements/animate"

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
)

func Build(uri string, data js.Object) js.Object {
	return js.Object{
		"uri":  uri,
		"data": data,
	}
}

func Chat(w vii.JsonWriter, username, channel, message string) {
	w.WriteJson(Build("chat", js.Object{
		"username": username,
		"channel":  channel,
		"message":  message,
	}))
}

func ChatJoin(w vii.JsonWriter, username, channel string, messages []js.Object) {
	w.WriteJson(Build("/chat/join", js.Object{
		"username": username,
		"channel":  channel,
		"messages": messages,
	}))
}

func Error(w vii.JsonWriter, source, message string) {
	w.WriteJson(Build("/error", js.Object{
		"source":  source,
		"message": message,
	}))
}
