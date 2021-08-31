package runtime

// // LoginSocket provides the login mechanism from the websocket connection
// func (t *T) LoginSocket(socket *websocket.T, account *account.T) (player *player.T, error error) {
// 	p, err := t.Players.Login(account)
// 	if err != nil {
// 		error = err
// 	} else {
// 		p.AddConn(socket.ID())
// 		go t.PlayerSocketWaiter(player, socket)
// 		socket.Send("/data/myaccount", player.Account.JSON()) // RAW URL SEND
// 	}
// 	return
// }
