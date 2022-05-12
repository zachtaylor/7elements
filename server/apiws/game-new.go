package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/match"
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func GameNew(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("no session")
			socket.Write(wsout.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", socket.SessionID())

		user, _, err := server.GetUserManager().GetSession(socket.SessionID())
		if user == nil {
			log.Add("Error", err).Error("user missing")
			socket.Write(wsout.Error("vii", "internal error"))
			return
		}
		log = log.Add("Username", user.Name())

		account := server.GetAccounts().Get(user.Name())
		if account == nil {
			log.Error("account missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		} else if account.GameID != "" {
			log.Add("Game", account.GameID).Warn("game exists")
			socket.WriteSync(wsout.Error("vii", "game exists"))
			return
		}

		if queue := server.GetMatchMaker().Get(user.Name()); queue != nil {
			log.Add("QueueStart", queue.Start().Format(`15:04:05`)).Warn("queue exists")
			socket.WriteSync(wsout.Error("vii", "already searching for game"))
			return
		}

		version := server.GetGameVersion()

		var deckid int
		if deckid64, _ := m.Data["deckid"].(float64); deckid64 < 1 {
			log.Add("Data", m.Data).Warn("deckid missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		} else {
			deckid = int(deckid64)
		}

		var owner string
		if owner, _ = m.Data["owner"].(string); owner == "" {
			log.Add("Data", m.Data).Warn("owner missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		}

		pvp := false
		if _pvp, ok := m.Data["pvp"].(bool); ok && _pvp {
			pvp = true
		} else if ok {
			pvp = false
		} else {
			log.Add("val", m.Data["pvp"]).Add("type", reflect.TypeOf(m.Data["pvp"])).Warn("pvp missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		}

		var hands string
		if hands = m.Data["hands"].(string); hands == "" {
			log.Add("Data", m.Data).Warn("hands missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		} else if hands == "small" || hands == "med" || hands == "large" {
			// ok
		} else {
			log.Add("Data", m.Data).Warn("unknown hands")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		}

		var speed string
		if speed = m.Data["speed"].(string); speed == "" {
			log.Add("Data", m.Data).Warn("speed missing")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		} else if speed == "slow" || speed == "med" || speed == "fast" {
			// ok
		} else {
			log.Add("Data", m.Data).Warn("unknown speed")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		}

		var decklist *deck.Prototype
		if owner == "vii" {
			decklist = version.GetDecks()[deckid]
			if decklist == nil {
				log.Warn("deckid invalid")
				socket.WriteSync(wsout.Error("vii", "internal error"))
				return
			}
		} else if owner != account.Username {
			log.Add("Owner", owner).Warn("owner unexpected")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		} else {
			decklist = account.Decks[deckid]
			if decklist == nil {
				log.Warn("deckid invalid")
				socket.WriteSync(wsout.Error("vii", "internal error"))
				return
			} else if decklist.Count() < 21 {
				log.Warn("deck too small")
				socket.WriteSync(wsout.Error("vii", "deck must have at least 21 cards"))
				return
			}
			log.Add("Deck", decklist.ID)
		}

		settings := match.NewQueueSettings(decklist, hands, speed)

		log.Trace("ready")

		// find game
		var game *game.T
		if !pvp {
			log.Info("vs ai")
			ai := ai.New("A.I.")
			game = server.GetGameManager().New(settings.Rules(), game.NewEntry(decklist, user), ai.Entry(version))
			ai.Connect(game)
		} else if q, err := server.GetMatchMaker().Queue(user, settings); q == nil {
			log.Add("Error", err).Error("queue error")
			socket.WriteSync(wsout.Error("vii", "internal error"))
			return
		} else {
			log.Info("wait")
			socket.WriteSync(wsout.Queue(q.Data()))
			if gameid := q.SyncGameID(); len(gameid) < 1 {
				log.Add("Error", err).Error("queue error")
				socket.WriteSync(wsout.Error("vii", "internal error"))
				return
			} else if game = server.GetGameManager().Get(gameid); game == nil {
				log.Add("Error", err).Error("game error")
				socket.WriteSync(wsout.Error("vii", "internal error"))
				return
			}
			log.Info("done")
		}

		account.GameID = game.ID()
		socket.Write(wsout.MyAccount(account.Data()).EncodeToJSON())
	})
}
