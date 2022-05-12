package request

// import (
// 	"github.com/zachtaylor/7elements/game"
// 	"github.com/zachtaylor/7elements/game/seat"
// 	"taylz.io/http/websocket"
// )

// func Chat(game *game.T, seat *seat.T, json map[string]any) {
// 	text, _ := json["text"].(string)
// 	game.Log().With(map[string]any{
// 		"Username": seat.Username,
// 		"Text":     text,
// 	}).Trace()
// 	go game.Chat(seat.Username, text)
// }
