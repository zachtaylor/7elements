package seat

// import (
// 	"strconv"

// 	"github.com/zachtaylor/7elements/deck"
// 	"github.com/zachtaylor/7elements/element"
// 	"taylz.io/yas"
// )

// type T struct {
// 	Username string
// 	Life     int
// 	Future   *deck.T
// 	Karma    element.Karma
// 	Hand     yas.Set[string]
// 	Present  yas.Set[string]
// 	Past     yas.Set[string]
// 	Color    string
// 	Writer   Writer
// }

// func New(life int, deck *deck.T, writer Writer) *T {
// 	return &T{
// 		Username: deck.User,
// 		Life:     life,
// 		Future:   deck,
// 		Karma:    element.Karma{},
// 		Hand:     yas.Set[string]{},
// 		Present:  yas.Set[string]{},
// 		Past:     yas.Set[string]{},
// 		Writer:   writer,
// 	}
// }

// // Message sends data to player agent if available
// // func (seat *T) Message(uri string, json map[string]any) {
// // 	if seat.Writer != nil {
// // 		seat.Writer.Write(websocket.Message{URI: uri, Data: json}.EncodeToJSON())
// // 	}
// // }

// func (seat *T) String() string {
// 	if seat == nil {
// 		return "<nil>"
// 	}
// 	return `{` +
// 		seat.Username +
// 		` ♥:` + strconv.FormatInt(int64(seat.Life), 10) +
// 		" " + seat.Karma.String() +
// 		` ◘:` + strconv.FormatInt(int64(len(seat.Hand)), 10) +
// 		`}`
// }

// // JSON returns representation of a game seat as map[string]any
// func (seat *T) JSON() map[string]any {
// 	return map[string]any{
// 		"username": seat.Username,
// 		"future":   len(seat.Future.Cards),
// 		"life":     seat.Life,
// 		"present":  seat.Present.Slice(),
// 		"hand":     len(seat.Hand),
// 		"elements": seat.Karma.JSON(),
// 		"past":     seat.Past.Slice(),
// 	}
// }
