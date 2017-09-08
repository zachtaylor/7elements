package log

func new() Log {
	return Log{}
}

func Add(name string, value interface{}) Log {
	return new().Add(name, value)
}

func Debug(message string) {
	new().Debug(message)
}

func Info(message string) {
	new().Info(message)
}

func Warn(message string) {
	new().Warn(message)
}

func Error(message interface{}) {
	new().Error(message)
}
