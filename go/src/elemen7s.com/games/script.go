package games

var Scripts = make(map[string]Script)

type Script func(*Game, *Seat, interface{})
