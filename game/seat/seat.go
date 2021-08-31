package seat

import (
	"strconv"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game/token"
	"taylz.io/http/websocket"
)

type T struct {
	Username string
	Life     int
	Deck     *deck.T
	Karma    element.Karma
	Hand     card.Set
	Present  token.Map
	Past     card.Set
	Color    string
	Writer   websocket.Writer
}

func New(life int, deck *deck.T, writer websocket.Writer) *T {
	return &T{
		Username: deck.User,
		Life:     life,
		Deck:     deck,
		Karma:    element.Karma{},
		Hand:     card.Set{},
		Present:  token.Map{},
		Past:     card.Set{},
		Writer:   writer,
	}
}

// Message sends data to player agent if available
func (seat *T) Message(uri string, json websocket.MsgData) {
	if seat.Writer != nil {
		seat.Writer.Message(uri, json)
	}
}

func (seat *T) String() string {
	if seat == nil {
		return "<nil>"
	}
	return `{` + seat.Username +
		` ♥:` + strconv.FormatInt(int64(seat.Life), 10) +
		` ☼:` + seat.Karma.String() +
		` ♣:` + strconv.FormatInt(int64(len(seat.Hand)), 10) +
		` ◘:` + strconv.FormatInt(int64(len(seat.Deck.Cards)), 10) +
		`}`
}

// Data returns JSON representation of a game seat
func (seat *T) Data() websocket.MsgData {
	return websocket.MsgData{
		"username": seat.Username,
		"deck":     len(seat.Deck.Cards),
		"life":     seat.Life,
		"present":  seat.Present.Keys(),
		"hand":     len(seat.Hand),
		"elements": seat.Karma.JSON(),
		"past":     seat.Past.Keys(),
		"future":   len(seat.Deck.Cards),
	}
}
