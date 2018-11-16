package animate // import "github.com/zachtaylor/7elements/animate"

// Build creates a push request object
//
// This macro creates js.Object{"uri":uri,"data":data}
func Build(uri string, data Json) Json {
	return Json{
		"uri":  uri,
		"data": data,
	}
}

// Alert creates an alert push request object with Build
//
// alert level should be in ["info", "warn", "error"]
func Alert(class, source, text string) Json {
	return Build("/alert", Json{
		"level":  class,
		"source": source,
		"text":   text,
	})
}

// AlertInfo calls Alert with "info"
func AlertInfo(source, text string) Json {
	return Alert("info", source, text)
}

// AlertWarn calls Alert with "warn"
func AlertWarn(source, text string) Json {
	return Alert("warn", source, text)
}

// AlertError calls Alert with "error"
func AlertError(source, text string) Json {
	return Alert("error", source, text)
}

func Chat(w JsonWriter, username, channel, message string) {
	w.WriteJson(Build("/chat", Json{
		"username": username,
		"channel":  channel,
		"message":  message,
	}))
}

func ChatJoin(w JsonWriter, username, channel string, messages []Json) {
	w.WriteJson(Build("/chat/join", Json{
		"username": username,
		"channel":  channel,
		"messages": messages,
	}))
}

func Error(w JsonWriter, source, message string) {
	w.WriteJson(Build("/error", Json{
		"source":  source,
		"message": message,
	}))
}
