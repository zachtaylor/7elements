package accountscards

import (
	"fmt"
	"time"
	"ztaylor.me/js"
)

type AccountCard struct {
	Username string
	CardId   int
	Register time.Time
	Notes    string
}

type Stack map[int][]*AccountCard

func (stack Stack) Json() js.Object {
	j := js.Object{}
	for cardId, list := range stack {
		j[fmt.Sprintf("%d", cardId)] = len(list)
	}
	return j
}
