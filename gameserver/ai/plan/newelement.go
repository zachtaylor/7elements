package plan

import (
	"strconv"

	"github.com/zachtaylor/7elements/element"
	"taylz.io/http/websocket"
)

// NewElement is a plan to choose an element
type NewElement struct {
	StateID string
	Element element.T
}

func (el *NewElement) Score() int {
	return 12
}

func (el *NewElement) Submit(request RequestFunc) {
	request(el.StateID, websocket.MsgData{
		"choice": strconv.FormatInt(int64(el.Element), 10),
	})
}

func (el *NewElement) String() string {
	return "Choice Element"
}
