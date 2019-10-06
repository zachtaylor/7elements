package apiws

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func GameNew(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Root.Logger.New().Tag("apiws/game.new").With(log.Fields{
			"Username": m.User,
		})

		if socket.Session == nil {
			log.Warn("session required")
			return
		} else if game := rt.Games.FindUsername(m.User); game != nil {
			log.Add("GameID", game.ID).Warn("game exists")
			connectGame(rt, log, socket)
			return
		}

		var deck *vii.AccountDeck
		if deckid := m.Data.GetI("deckid"); deckid < 1 {
			log.Add("DeckID", m.Data["deckid"]).Warn("deckid parse error")
			return
		} else if account, ok := m.Data["account"]; !ok {
			log.Warn("account missing")
			return
		} else if cast.Bool(account) {
			deck = api.GetMyDeck(rt, log, m.User, deckid)
		} else {
			if deck = api.GetFreeDeck(rt, log, m.User, deckid); deck == nil {
				return
			}
		}
		log.Add("UseP2P", m.Data["account"]).Add("DeckID", m.Data["deckid"])

		if deck == nil {
			log.Warn("deck missing")
			return
		} else if deck.Count() < 20 {
			log.Warn("deck too small")
			pushJSON(socket, "/error", cast.JSON{
				"error": "deck must have at least 21 cards",
			})
			return
		}

		// build opponent
		// todo: bool(Query().Get("ai")) could be Query().Get("opponent") == "ai"

		var g *game.T
		if useai, ok := m.Data["ai"]; !ok {
			log.Warn("ai missing")
			return
		} else if cast.Bool(useai) {
			g = rt.Games.New(deck, ai.GetAccountDeck(rt.Root.Decks))
			ai.ConnectAI(g)
			log.Add("GameID", g.ID()).Info("created game vs ai")
		} else if search := rt.Games.Search(deck); search == nil {
			log.Warn("cannot start search")
		} else {
			log.Info("starting search")
			if gameid := <-search.Done; gameid == "" {
				log.Warn("search failed")
			} else {
				g = rt.Games.Get(gameid)
				log.Info("match found")
			}
		}

		if g == nil {
			log.Error("fail")
		} else {
			connectGame(rt, log, socket)
		}
	})
}
