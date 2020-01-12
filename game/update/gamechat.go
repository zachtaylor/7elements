package update

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
)

func GameChat(g *game.T, source, message string) {
	go g.GetChat().AddMessage(chat.NewMessage(source, message))
}
