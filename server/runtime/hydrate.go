package runtime

import (
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func (rt *T) hydrate(ws *websocket.T) {
	log := rt.Logger.Add("Socket", ws.ID())
	user, session, err := rt.Users.GetSession(ws.SessionID())
	if err != nil {
		log.Error(err)
		return
	}
	session.Update()
	log = log.Add("SessionID", session.ID()).Add("Username", user.Name())
	account := rt.Accounts.Get(user.Name())
	if account == nil {
		log.Warn("account missing")
		return
	}

	ws.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())

	if account.GameID == "" {
	} else if game := rt.Games.Get(account.GameID); game == nil {
		log.Add("GameID", account.GameID).Warn("game expired")
		account.GameID = ""
	} else if seat := game.Seats.Get(account.Username); seat == nil {
		log.Add("Seats", game.Seats.Keys()).Add("GameID", game.ID()).Warn("seat expired")
		account.GameID = ""
	} else {
		game.Request(account.Username, "connect", map[string]interface{}{})
	}

	if queue := rt.MatchMaker.Get(account.Username); queue != nil {
		ws.WriteSync(wsout.Queue(queue.Data()))
	}
}
