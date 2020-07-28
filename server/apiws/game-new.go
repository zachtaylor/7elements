package apiws

import (
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func GameNew(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Log().Tag("apiws/game.new").With(cast.JSON{
			"Session": socket.Session,
		})

		if socket.Session == nil {
			log.Warn("session required")
			return
		} else if game := rt.Games.FindUsername(socket.Session.Name()); game != nil {
			log.Add("GameID", game.ID).Warn("game exists")
			connectgame(rt, socket)
			return
		}

		var decklist *deck.Prototype
		if deckid := m.Data.GetI("deckid"); deckid < 1 {
			log.Add("DeckID", m.Data["deckid"]).Warn("deckid parse error")
			return
		} else if list, err := rt.Decks.Get(deckid); err != nil {
			log.Add("Error", err).Error("account missing")
			return
		} else if list.User != "vii" && list.User != socket.Session.Name() {
			log.Add("Owner", list.User).Warn("deckid is not public")
			return
		} else if list.Count() < 21 {
			log.Warn("deck too small")
			out.Error(socket, list.Name, "deck must have at least 21 cards")
			return
		} else {
			log.Add("DeckID", deckid)
			decklist = list
		}
		deck := deck.New(rt.Logger, rt.Cards, decklist, socket.Session.Name())
		if deck == nil {
			log.Warn("failed to instantiate")
			out.Error(socket, "vii", "internal server error")
			return
		}

		// build opponent
		// todo: bool(Query().Get("ai")) could be Query().Get("opponent") == "ai"

		var g *game.T
		if useai, ok := m.Data["ai"]; !ok {
			log.Warn("ai missing")
			return
		} else if cast.Bool(useai) {
			g = rt.Games.New(deck, ai.GetDeck(rt.Logger, rt.Cards, rt.Decks))
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
			connectgame(rt, socket)
		}
	})
}
