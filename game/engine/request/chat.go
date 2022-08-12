package request

// import (
// 	"github.com/zachtaylor/7elements/game"
// 	"github.com/zachtaylor/7elements/game/seat"
// 	"taylz.io/http/websocket"
// )

// func Chat(game *game.G, player *game.Player, json map[string]any) {
// 	text, _ := json["text"].(string)
// 	game.Log().With(map[string]any{
// 		"Username": player.T.Writer.Name(),
// 		"Text":     text,
// 	}).Trace()
// 	go game.Chat(player.T.Writer.Name(), text)
// }
