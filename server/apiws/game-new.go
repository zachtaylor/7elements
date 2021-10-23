package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/match"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func GameNew(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("no session")
			socket.Write(wsout.ErrorJSON("vii", "no user"))
			return
		}
		log = log.Add("Session", socket.SessionID())

		user, _, err := rt.Users.GetSession(socket.SessionID())
		if user == nil {
			log.Add("Error", err).Error("user missing")
			socket.Write(wsout.ErrorJSON("vii", "internal error"))
			return
		}
		log = log.Add("Username", user.Name())

		account := rt.Accounts.Get(user.Name())
		if account == nil {
			log.Error("account missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else if account.GameID != "" {
			log.Add("Game", account.GameID).Warn("game exists")
			socket.WriteSync(wsout.ErrorJSON("vii", "game exists"))
			return
		}

		if queue := rt.MatchMaker.Get(user.Name()); queue != nil {
			log.Add("QueueStart", queue.Start().Format(`15:04:05`)).Warn("queue exists")
			socket.WriteSync(wsout.ErrorJSON("vii", "already searching for game"))
			return
		}

		var deckid int
		if deckid64, _ := m.Data["deckid"].(float64); deckid64 < 1 {
			log.Add("Data", m.Data).Warn("deckid missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else {
			deckid = int(deckid64)
		}

		var owner string
		if owner, _ = m.Data["owner"].(string); owner == "" {
			log.Add("Data", m.Data).Warn("owner missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		pvp := false
		if _pvp, ok := m.Data["pvp"].(bool); ok && _pvp {
			pvp = true
		} else if ok {
			pvp = false
		} else {
			log.Add("val", m.Data["pvp"]).Add("type", reflect.TypeOf(m.Data["pvp"])).Warn("pvp missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		var hands string
		if hands = m.Data["hands"].(string); hands == "" {
			log.Add("Data", m.Data).Warn("hands missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else if hands == "small" || hands == "med" || hands == "large" {
			// ok
		} else {
			log.Add("Data", m.Data).Warn("unknown hands")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		var speed string
		if speed = m.Data["speed"].(string); speed == "" {
			log.Add("Data", m.Data).Warn("speed missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else if speed == "slow" || speed == "med" || speed == "fast" {
			// ok
		} else {
			log.Add("Data", m.Data).Warn("unknown speed")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		var decklist *deck.Prototype
		if owner == "vii" {
			decklist = rt.Decks[deckid]
			if decklist == nil {
				log.Warn("deckid invalid")
				socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
				return
			}
		} else if owner != account.Username {
			log.Add("Owner", owner).Warn("owner unexpected")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else {
			decklist = account.Decks[deckid]
			if decklist == nil {
				log.Warn("deckid invalid")
				socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
				return
			} else if decklist.Count() < 21 {
				log.Warn("deck too small")
				socket.WriteSync(wsout.ErrorJSON("vii", "deck must have at least 21 cards"))
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
			game = rt.Games.New(settings.Rules(), game.NewEntry(decklist, user), ai.Entry(rt.Logger, rt.Cards, rt.Decks))
			ai.Connect(game)
		} else if q, err := rt.MatchMaker.Queue(user, settings); q == nil {
			log.Add("Error", err).Error("queue error")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else {
			log.Info("wait")
			socket.WriteSync(wsout.Queue(q.Data()))
			if gameid := q.SyncGameID(); len(gameid) < 1 {
				log.Add("Error", err).Error("queue error")
				socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
				return
			} else if game = rt.Games.Get(gameid); game == nil {
				log.Add("Error", err).Error("game error")
				socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
				return
			}
			log.Info("done")
		}

		account.GameID = game.ID()
		socket.Write(wsout.MyAccount(account.Data()).EncodeToJSON())
	})
}
