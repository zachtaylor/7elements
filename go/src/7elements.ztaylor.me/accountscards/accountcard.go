package accountscards

import (
	"time"
	"ztaylor.me/json"
)

type AccountCard struct {
	Username string
	CardId   int
	Register time.Time
	Notes    string
}

type Stack map[int][]*AccountCard

func (stack Stack) Json() json.Json {
	j := json.Json{}
	for cardId, list := range stack {
		j[json.UItoS(uint(cardId))] = len(list)
	}
	return j
}
