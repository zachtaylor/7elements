package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/gameserver"
	"github.com/zachtaylor/7elements/gameserver/ai"
	"github.com/zachtaylor/7elements/gameserver/queue"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
	"taylz.io/types"
)

func GameNew(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Logger.Add("Socket", socket.ID())
		if len(socket.Name()) < 1 {
			log.Warn("anon update email")
			socket.WriteSync(wsout.ErrorJSON("vii", "you must log in to change email"))
			return
		}
		log.Add("Name", socket.Name())

		user := rt.Users.Get(socket.Name())
		if user == nil {
			log.Error("user missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		account := rt.Accounts.Get(socket.Name())
		if account == nil {
			log.Error("user missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else if account.GameID != "" {
			log.Add("Game", account.GameID).Warn("game exists")
			socket.WriteSync(wsout.ErrorJSON("vii", "game exists"))
			return
		}

		if queue := rt.Queue.Get(socket.Name()); queue != nil {
			log.Add("Queue", queue).Warn("queue exists")
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

		owner, _ := m.Data["owner"].(string)
		if owner == "" {
			log.Add("Data", m.Data).Warn("owner missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		// search settings
		var options queue.Options

		if custom, ok := m.Data["custom"].(bool); ok && custom {
			options.AllowCustomDecks = true
		} else if ok {
			// options.AllowCustomDecks = false
		} else {
			log.Add("val", m.Data["custom"]).Add("type", reflect.TypeOf(m.Data["custom"])).Warn("custom missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		if speed := m.Data["speed"]; speed == "fast" {
			options.Rules.Timeout = 30 * types.Second
		} else if speed == "med" {
			options.Rules.Timeout = 60 * types.Second
		} else if speed == "slow" {
			options.Rules.Timeout = 90 * types.Second
		} else {
			log.Add("Data", m.Data).Warn("speed missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		if handsize := m.Data["hands"]; handsize == "small" {
			options.Rules.StartingHand = 3
		} else if handsize == "med" {
			options.Rules.StartingHand = 5
		} else if handsize == "large" {
			options.Rules.StartingHand = 7
		} else {
			log.Add("Data", m.Data).Warn("handsize missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		}

		var decklist *deck.Prototype
		if owner == "vii" {
			decklist = rt.Decks[deckid]
			if decklist == nil {
				log.Add("Owner", owner).Warn("deckid invalid")
				socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
				return
			}
			log.Add("Deck", decklist.Name)
		} else if owner != account.Username {
			log.Add("Owner", owner).Warn("owner unexpected")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else {
			options.AllowCustomDecks = true
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

		// find game

		deck := deck.New(rt.Logger, rt.Cards, decklist, account.Username)
		entry := &gameserver.Entry{
			Deck:   deck,
			Writer: user,
		}

		var game *game.T

		if pvp, ok := m.Data["pvp"].(bool); ok && !pvp {
			ai := ai.New("A.I.")
			game = rt.Games.New(options.Rules, rt.Logger, entry, ai.Entry(rt.Logger, rt.Cards, rt.Decks))
			ai.Connect(game)
			log.Info("created game vs ai")
		} else if !ok {
			log.Add("pvp", m.Data["pvp"]).Add("type", reflect.TypeOf(m.Data["pvp"])).Warn("pvp missing")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else if queue, err := rt.Queue.Request(options, user, deck); err != nil {
			log.Add("Error", err).Error("queue error")
			socket.WriteSync(wsout.ErrorJSON("vii", "internal error"))
			return
		} else {
			select {
			case id := <-queue.Done():
				game = rt.Games.Get(id)
			case <-queue.Cancel():
				log.Trace("queue cancelled")
				return
			}
		}

		if game == nil {
			log.Error("fail")
		} else {
			log.Info("finish")
			account.GameID = game.ID()
			socket.Write(wsout.MyAccountGame(game.ID()).EncodeToJSON())
		}
	})
}
