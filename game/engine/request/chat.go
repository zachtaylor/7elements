package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"taylz.io/http/websocket"
)

func Chat(game *game.T, seat *seat.T, json websocket.MsgData) {
	text, _ := json["text"].(string)
	game.Log().With(websocket.MsgData{
		"Username": seat.Username,
		"Text":     text,
	}).Trace()
	go game.Chat(seat.Username, text)
}
