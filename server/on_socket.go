package server

import (
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

// OnSocket creates a closure for the internal.Server definition to react to a change in socket cache
func OnSocket(server internal.Server) websocket.CacheObserver {
	return func(id string, oldWS, newWS *websocket.T) {
		go OnSocketServer(server, id, oldWS, newWS) // surrender hotpath instantly
	}
}

func OnSocketServer(server internal.Server, id string, oldWS, newWS *websocket.T) {
	log := server.Log().Add("Socket", id)
	if oldWS == nil && newWS != nil {

		if len(newWS.SessionID()) > 0 {
			log = log.Add("SessionID", newWS.SessionID())
			OnSocketHydrate(server, newWS)
		}
		log.Trace("open")
	} else if oldWS != nil && newWS == nil {
		if len(oldWS.SessionID()) > 0 {
			log = log.Add("SessionID", oldWS.SessionID())
		}
		log.Trace("close")
	} else {
		log.Add("Old", oldWS).Add("New", newWS).Warn("weird")
	}
	server.Ping()
}

func OnSocketHydrate(server internal.Server, ws *websocket.T) {
	log := server.Log().Add("Socket", ws.ID())
	user, session, err := server.GetUserManager().GetSession(ws.SessionID())
	if err != nil {
		log.Error(err)
		return
	}
	session.Update()
	log = log.Add("SessionID", session.ID()).Add("Username", user.Name())
	account := server.GetAccounts().Get(user.Name())
	if account == nil {
		log.Warn("account missing")
		return
	}

	ws.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())

	if account.GameID == "" {
	} else if game := server.GetGameManager().Get(account.GameID); game == nil {
		log.Add("GameID", account.GameID).Warn("game expired")
		account.GameID = ""
	} else if seat := game.Seats.Get(account.Username); seat == nil {
		log.Add("Seats", game.Seats.Keys()).Add("GameID", game.ID()).Warn("seat expired")
		account.GameID = ""
	} else {
		game.Request(account.Username, "connect", map[string]any{})
	}

	if queue := server.GetMatchMaker().Get(account.Username); queue != nil {
		ws.WriteSync(wsout.Queue(queue.Data()))
	}
}
