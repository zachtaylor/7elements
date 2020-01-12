package apiws

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func _connectGameWaiter(socket *websocket.T, game *game.T, seat string, log *log.Entry) {
	select {
	case <-game.Done():
	case <-socket.DoneChan():
	case <-socket.Session.Done():
	}
	log.Source().Debug()
	game.GetSeat(seat).Receiver = nil
}
