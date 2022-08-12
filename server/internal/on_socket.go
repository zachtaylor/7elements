package internal

import (
	"github.com/zachtaylor/7elements/game"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/log"
)

// OnSocket creates a closure for the internal.Server definition to react to a change in socket cache
func OnSocket(server Server) websocket.ObserverFunc {
	return func(id string, newWS, oldWS *websocket.T) {
		go OnSocketServer(server, id, newWS, oldWS) // surrender hotpath instantly
	}
}

func OnSocketServer(server Server, id string, newWS, oldWS *websocket.T) {
	log := server.Log().Add("Socket", id)
	if oldWS == nil && newWS != nil {
		if user, err := server.Users().ReadHTTP(newWS.Request()); user == nil {
			log = log.Add("SessionErr", err)
		} else {
			log = log.Add("Session", user.Session().ID())
			OnSocketHydrate(server, user, newWS, log.New())
		}
		log.Trace("open")
	} else if oldWS != nil && newWS == nil {
		if user := server.Users().GetWebsocket(oldWS); user != nil {
			log = log.Add("Session", user.Session().ID())
		}
		log.Trace("close")
	} else {
		log.Add("Old", oldWS).Add("New", newWS).Warn("weird")
	}
	server.Ping()
}

func OnSocketHydrate(server Server, user *user.T, ws *websocket.T, log log.Writer) {
	server.Sessions().Update(user.Session().ID())

	account := server.Accounts().Get(user.Session().Name())
	if account == nil {
		log.Warn("account missing")
		return
	}

	ws.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))

	if account.GameID == "" || account.PlayerID == "" {
		log.Debug("no game")
	} else if g := server.Games().Get(account.GameID); g == nil {
		log.Add("GameID", account.GameID).Warn("game expired")
		account.GameID, account.PlayerID = "", ""
	} else if player := g.Player(account.PlayerID); player == nil {
		log.Add("Seats", g.Players()).Add("GameID", g.ID()).Warn("seat expired", account.PlayerID)
		account.GameID, account.PlayerID = "", ""
	} else {
		// player.T.Writer =
		g.AddRequest(game.NewReq(account.Username, "connect", map[string]any{}))
	}

	if queue := server.MatchMaker().Get(account.Username); queue != nil {
		ws.WriteText(websocket.NewMessage("/game/queue", queue.Data()).ShouldMarshal())
	}
}
