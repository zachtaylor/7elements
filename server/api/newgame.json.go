package api

import (
	"net/http"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/game/engine"
	"ztaylor.me/cast"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func NewGameHandler(sessions *sessions.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := sessions.ReadRequestCookie(r)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.WithFields(log.Fields{
				"RemoteAddr": r.RemoteAddr,
			}).Warn("api/newgame: session required")
			return
		}

		account, err := vii.AccountService.Get(session.Name())
		if account == nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.WithFields(log.Fields{
				"RemoteAddr": r.RemoteAddr,
			}).Warn("api/newgame: account missing")
			return
		}

		var deck *vii.AccountDeck
		if _deckid := r.URL.Query().Get("deckid"); len(_deckid) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Query", r.URL.Query().Encode()).Warn("api/newgame: deckid missing")
			return
		} else if deckid := cast.Int(_deckid); deckid < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.WithFields(log.Fields{
				"DeckID": _deckid,
				"Error":  err,
			}).Warn("api/newgame: deckid missing")
			return
		} else if usep2p := r.URL.Query().Get("usep2p"); len(usep2p) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Error", err).Warn("api/newgame: usep2p missing")
			return
		} else if usep2p == "true" {
			if mydecks, err := vii.AccountDeckService.Get(session.Name()); mydecks == nil {
				log.Add("Error", err).Error("api/newgame: decks missing")
			} else if d := mydecks[deckid]; deck == nil {
				log.Add("Deck", deck).Error("api/newgame: start failed, deck missing")
			} else if d.Count() < 21 {
				// vii.NotificationService.Add(session.Name(), "error", "New Game", "Your deck must have at least 21 cards")
				log.WithFields(log.Fields{
					"Username": session.Name(),
					"UseP2P":   usep2p,
					"DeckID":   _deckid,
					"Count":    d.Count(),
					"Error":    err,
				}).Warn("api/newgame: start failed, deck too small")
			} else {
				deck = d
			}
		} else {
			if d, err := vii.DeckService.Get(deckid); d == nil {
				log.WithFields(log.Fields{
					"Username": session.Name(),
					"UseP2P":   usep2p,
					"DeckID":   deckid,
					"Error":    err,
				}).Error("api/newgame: deck missing")
			} else {
				deck = vii.NewAccountDeckWith(d, session.Name())
			}
		}

		// todo: bool(Query().Get("ai")) -> Query().Get("op") == "ai"

		var g *game.T
		if _ai := r.URL.Query().Get("ai"); len(_ai) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Error", err).Warn("api/newgame: ai missing")
			return
		} else if useai := cast.Bool(_ai); useai {
			g = game.Service.New()
			g.Register(deck)
			ai.Register(g)
			engine.Watch(g)
			log.WithFields(log.Fields{
				"GameID":   g.ID(),
				"Username": session.Name,
			}).Info("api/newgame: created game vs ai")
		} else if search := game.Service.StartPlayerSearch(deck); search == nil {
			log.WithFields(log.Fields{
				"Session": session,
				"DeckId":  deck.ID,
			}).Warn("api/newgame: failed")
		} else if gameid := <-search.Done; gameid == "" {
			log.WithFields(log.Fields{
				"Session": session,
				"GameId":  gameid,
			}).Warn("api/newgame: search failed")
		} else {
			g = game.Service.Get(gameid)
		}

		if g != nil {
			w.Write([]byte(g.Json(session.Name()).String()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
	})
}
