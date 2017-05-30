package log

func New() *Entries {
	return &Entries{}
}

func Clone() *Entries {
	clone := New()
	for k, v := range *globalEntries() {
		clone.Add(k, v)
	}
	return clone
}

func Add(name string, value interface{}) *Entries {
	return globalEntries().Add(name, value)
}

func Debug(message string) {
	globalEntries().Debug(message)
}

func Info(message string) {
	globalEntries().Info(message)
}

func Warn(message string) {
	globalEntries().Warn(message)
}

func Error(message interface{}) {
	globalEntries().Error(message)
}
