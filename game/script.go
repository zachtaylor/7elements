package game

type ScriptFunc = func(*T, *Seat, interface{}) []Event

var Scripts = make(map[string]ScriptFunc)
