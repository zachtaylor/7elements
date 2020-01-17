package request

import (
	pkg_chat "github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func chat(g *game.T, seat *game.Seat, json cast.JSON) {
	text := json.GetS("text")
	g.Log().With(cast.JSON{
		"Username": seat.Username,
		"Text":     text,
	}).Debug("engine/chat") // died after
	go g.GetChat().AddMessage(pkg_chat.NewMessage(seat.Username, text))
}
