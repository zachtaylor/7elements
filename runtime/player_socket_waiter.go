package runtime

import (
	"github.com/zachtaylor/7elements/player"
	"ztaylor.me/http/websocket"
)

func (t *T) PlayerSocketWaiter(player *player.T, socket *websocket.T) {
	<-socket.DoneChan()
	player.RemConn(socket.ID())
}
