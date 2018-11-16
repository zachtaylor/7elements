package api

import (
	"net/http"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/engine/ai"
	"ztaylor.me/cast"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func NewGameHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := sessions.ReadCookie(r)
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
		if _deckid, ok := r.URL.Query()["deckid"]; !ok || len(_deckid) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Error", err).Warn("api/newgame: deckid missing")
			return
		} else if deckid := cast.Int(_deckid[0]); deckid < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.WithFields(log.Fields{
				"DeckID": _deckid,
				"Error":  err,
			}).Warn("api/newgame: deckid missing")
			return
		} else if _usep2p, ok := r.URL.Query()["usep2p"]; !ok || len(_usep2p) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Error", err).Warn("api/newgame: usep2p missing")
			return
		} else if usep2p := _usep2p[0]; usep2p == "true" {
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

		var game *vii.Game
		if _ai, ok := r.URL.Query()["deckid"]; !ok || len(_ai) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Add("Error", err).Warn("api/newgame: ai missing")
			return
		} else if useai := cast.Bool(_ai[0]); useai {
			game = vii.GameService.New()
			game.Register(deck)
			ai.Register(game)
			vii.GameService.Watch(game)
			log.WithFields(log.Fields{
				"GameId":   game,
				"Username": session.Name,
			}).Info("api/newgame: created game vs ai")
		} else if search := vii.GameService.StartPlayerSearch(deck); search == nil {
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
			game = vii.GameService.Get(gameid)
		}

		if game != nil {
			w.Write([]byte(game.Json(session.Name()).String()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
	})
}
