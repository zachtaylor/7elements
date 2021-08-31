package runtime

import (
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func (t *T) OnWebsocket(id string, oldWS, newWS *websocket.T) {
	go t.onWebsocket(id, oldWS, newWS)
}

func (t *T) onWebsocket(id string, oldWS, newWS *websocket.T) {
	log := t.Logger.Add("Socket", id)
	t.Ping()
	if newWS == nil {
		log.Add("Name", oldWS.Name()).Trace("close")
	} else if session := t.Sessions.GetName(newWS.Name()); session == nil {
		log.Trace("open")
	} else if account := t.Accounts.Get(session.Name()); account == nil {
		log.Warn("open account missing")
	} else {
		log.Add("Account", account.Username).Add("Game", account.GameID).Trace("open")
		session.Update()

		if account.GameID == "" {
		} else if game := t.Games.Get(account.GameID); game == nil {
			log.Add("Game", account.GameID).Trace("game expired")
			account.GameID = ""
		} else if seat := game.Seats.Get(account.Username); seat == nil {
			log.Add("Seats", game.Seats.Keys()).Add("Game", account.GameID).Warn("seat expired")
			account.GameID = ""
		} else {
			game.Request(account.Username, "connect", map[string]interface{}{})
		}

		newWS.Write(wsout.MyAccount(account.Data()).EncodeToJSON())
	}
}
