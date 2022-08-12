package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/match"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func GameNew(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user == nil {
			log.Warn("no session")
			socket.Write(websocket.MessageText, out.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", user.Session().ID())

		account := server.Accounts().Get(user.Session().Name())
		if account == nil {
			log.Error("account missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		} else if account.GameID != "" {
			log.Add("Game", account.GameID).Warn("game exists")
			socket.Write(websocket.MessageText, out.Error("vii", "game exists"))
			return
		}

		if queue := server.MatchMaker().Get(user.Session().Name()); queue != nil {
			log.Add("QueueStart", queue.Start().Format(`15:04:05`)).Warn("queue exists")
			socket.Write(websocket.MessageText, out.Error("vii", "already searching for game"))
			return
		}

		var deckid int
		if deckid64, _ := m.Data["deckid"].(float64); deckid64 < 1 {
			log.Add("Data", m.Data).Warn("deckid missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		} else {
			deckid = int(deckid64)
		}

		var owner string
		if owner, _ = m.Data["owner"].(string); owner == "" {
			log.Add("Data", m.Data).Warn("owner missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		} else if owner != "vii" && owner != account.Username {
			log.Add("Data", m.Data).Warn("owner must be self or system")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		if owner == "vii" {
			if deckid >= len(server.Content().Decks()) {
				log.Add("Data", m.Data).Warn("deckid too high")
				socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
				return
			}
		} else if deckid >= len(account.Decks) {
			log.Add("Data", m.Data).Warn("deckid too high")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		pvp := false
		if _pvp, ok := m.Data["pvp"].(bool); ok && _pvp {
			pvp = true
		} else if ok {
			pvp = false
		} else {
			log.Add("val", m.Data["pvp"]).Add("type", reflect.TypeOf(m.Data["pvp"])).Warn("pvp missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		var hands string
		if hands = m.Data["hands"].(string); hands == "" {
			log.Add("Data", m.Data).Warn("hands missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		} else if hands == "small" || hands == "med" || hands == "large" {
			// ok
		} else {
			log.Add("Data", m.Data).Warn("unknown hands")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		var speed string
		if speed = m.Data["speed"].(string); speed == "" {
			log.Add("Data", m.Data).Warn("speed missing")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		} else if speed == "slow" || speed == "med" || speed == "fast" {
			// ok
		} else {
			log.Add("Data", m.Data).Warn("unknown speed")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		}

		var cards card.Count
		// var decklist *deck.Prototype
		if owner == "vii" {
			cards = server.Content().Decks()[deckid].Cards
		} else {
			cards = account.Decks[deckid].Cards
		}

		settings := match.NewQueueSettings(owner, deckid, hands, speed)
		rules := match.RulesFromSettings(settings)
		if err := game.VerifyRulesDeck(rules, cards); err != nil {
			log.Add("Error", err).Debug("invalid deck")
			socket.Write(websocket.MessageText, out.Error("vii", err.Error()))
			return
		}

		log.Trace("ready")

		userWriter := game.UserWriter{User: user}

		// find game
		var g *game.G
		if !pvp {
			log.Info("vs ai")
			myEntry := game.NewEntry(userWriter, cards)
			g = server.MatchMaker().VSAI(myEntry, settings)
		} else if q, err := server.MatchMaker().Queue(userWriter, settings); q == nil {
			log.Add("Error", err).Error("queue error")
			socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
			return
		} else {
			log.Debug("waiting...")
			socket.WriteMessage(websocket.NewMessage("/game/queue", q.Data()))
			if gameid := q.Sync(); len(gameid) < 1 {
				log.Error("queue error")
				socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
				return
			} else if g = server.Games().Get(gameid); g == nil {
				log.Add("Error", err).Error("game error")
				socket.Write(websocket.MessageText, out.Error("vii", "internal error"))
				return
			}
			log.Info("done")
		}

		account.GameID = g.ID()
		socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
	})
}
