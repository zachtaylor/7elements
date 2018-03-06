package api

import (
	"bytes"
	"elemen7s.com/accountscards"
	"elemen7s.com/decks"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	zhttp "ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

var DecksIdJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	session, err := zhttp.ReadRequestCookie(r)
	log := log.Add("Addr", r.RemoteAddr)
	if err != nil || session == nil {
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("decks.id.json: session missing")
		return
	}

	log.Add("Username", session.Username)
	ds, err := decks.Get(session.Username)
	if err != nil {
		w.WriteHeader(500)
		log.Add("Error", err).Error("decks.json")
		return
	}

	var deckid = 0
	if deckidParse, err := strconv.Atoi(r.RequestURI[11 : len(r.RequestURI)-5]); err == nil {
		deckid = deckidParse
	} else {
		w.WriteHeader(500)
		w.Write([]byte("deckid parse error"))
		log.Add("Error", err).Error("decks.id.json: parse deck id")
		return
	}

	deck := ds[deckid]
	if deck == nil {
		w.WriteHeader(500)
		w.Write([]byte("deckid not found"))
		log.Add("Error", err).Error("decks.id.json: deck missing: " + session.Username)
		return
	}

	if r.Method == "GET" {
		deck.Json().Write(w)
		log.Debug("decks.id.json: get")
	} else if r.Method == "POST" {
		data := struct {
			Name  string
			Cards map[int]int
		}{}
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			log.Add("Error", err).Error("decks.id.json: post: data read")
			return
		} else if err := js.NewDecoder(bytes.NewBuffer(body)).Decode(&data); err != nil {
			log.Add("Body", string(body)).Add("Error", err).Error("decks.id.json: post: data parse")
			return
		}

		log.Add("DeckId", deckid)
		accountscards, err := accountscards.Get(session.Username)
		if accountscards == nil {
			log.Add("Error", err).Error("decks.id.json: post: cards missing")
			return
		}

		for cardid, count := range data.Cards {
			if maxCount := len(accountscards[cardid]); maxCount < count {
				log.Clone().Add("CardId", cardid).Add("RequestCount", count).Add("MaxCount", maxCount).Warn("decks.id.json: post: request more cards than in collection")
				deck.Cards[cardid] = maxCount
			} else if count == 0 {
				delete(deck.Cards, cardid)
			} else {
				deck.Cards[cardid] = count
			}
		}

		deck.Name = data.Name
		deck.Register = time.Now()

		if err := decks.Delete(session.Username, deck.Id); err != nil {
			log.Add("Error", err).Error("mydecks: post: delete old deck")
			return
		} else if err := decks.Insert(session.Username, deck.Id); err != nil {
			log.Add("Error", err).Error("mydecks: post: insert new deck")
			return
		}

		deck.Json().Write(w)
		log.Add("Name", deck.Name).Info("decks.id.json: post")
	}
})
