package runtime

import (
	"github.com/zachtaylor/7elements/player"
	"ztaylor.me/http/websocket"
)

// SignupSocket provides the signup mechanism from the websocket connection
func (t *T) SignupSocket(socket *websocket.T, username, email, password string) (player *player.T, error error) {
	p, err := t.Signup(username, email, password)
	if err != nil {
		error = err
	} else {
		p.AddConn(socket.ID())
		go t.PlayerSocketWaiter(player, socket)
		socket.Send("/data/myaccount", player.Account.JSON()) // RAW URL SEND
	}
	return
}
