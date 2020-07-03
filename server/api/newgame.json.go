package api

import (
	"net/http"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/server/runtime"
	"ztaylor.me/cast"
)

func NewGameHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log().Tag("api/newgame")
		session, _ := rt.Sessions.Cookie(r)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Add("RemoteAddr", r.RemoteAddr).Warn("session required")
			return
		}

		player := rt.Settings.Players.Get(session.Name())
		if player == nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Add("Error", err).Add("RemoteAddr", r.RemoteAddr).Warn("player missing")
		}

		log.Add("Username", player.Name())

		var deck *deck.T
		if _deckid := r.URL.Query().Get("deckid"); len(_deckid) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Query", r.URL.Query().Encode()).Warn("deckid missing")
			return
		} else if deckid := cast.Int(_deckid); deckid < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("DeckID", _deckid).Warn("deckid parse error")
			return
		} else if d := player; d != nil {

		} else {

		}

		// build opponent
		// todo: bool(Query().Get("ai")) could be Query().Get("opponent") == "ai"

		var g *game.T
		if _ai := r.URL.Query().Get("ai"); len(_ai) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Query", r.URL.Query().Encode()).Warn("ai missing")
			return
		} else if useai := cast.Bool(_ai); useai {
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

		if g != nil {
			w.Write(cast.BytesS(g.JSON(g.GetSeat(session.Name())).String()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
	})
}
