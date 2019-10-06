package api

import (
	"net/http"

	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"ztaylor.me/cast"
	"ztaylor.me/http/json"
)

func NewGameHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Tag("api/newgame")
		session := rt.Sessions.Cookie(r)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Add("RemoteAddr", r.RemoteAddr).Warn("session required")
			return
		}

		account, err := rt.Root.Accounts.Get(session.Name())
		if account == nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Add("Error", err).Add("RemoteAddr", r.RemoteAddr).Warn("account missing")
			return
		}

		log.Add("Username", session.Name())

		var deck *vii.AccountDeck
		if _deckid := r.URL.Query().Get("deckid"); len(_deckid) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Query", r.URL.Query().Encode()).Warn("deckid missing")
			return
		} else if deckid := cast.Int(_deckid); deckid < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("DeckID", _deckid).Warn("deckid parse error")
			return
		} else if usep2p := r.URL.Query().Get("usep2p"); len(usep2p) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Query", r.URL.Query().Encode()).Warn("usep2p missing")
			return
		} else if usep2p == "true" {
			if deck = GetMyDeck(rt, log, session.Name(), deckid); deck == nil {
				return
			}
		} else {
			if deck = GetFreeDeck(rt, log, session.Name(), deckid); deck == nil {
				return
			}
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
			w.Write(json.Encode(g.PerspectiveJSON(session.Name())))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
	})
}
