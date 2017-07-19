package mydecks

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	// "7elements.ztaylor.me/server/sessionman"
	// "7elements.ztaylor.me/server/util"
	// "bufio"
	"net/http"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request, deckid int, account *SE.Account) {
	log.Add("Username", account.Username)

	decoder := json.NewDecoder(r.Body)
	deck := SE.NewAccountDeck()
	err := decoder.Decode(deck)
	if err != nil {
		log.Add("Error", err).Error("mydecks: post: data parse")
		return
	}

	accountsdecks := SE.AccountsDecks.Cache[account.Username]
	if accountsdecks == nil {
		log.Error("mydecks: post: accountsdecks missing")
		return
	}
	if accountsdecks[deck.Id] == nil {
		log.Add("DeckId", deck.Id).Error("mydecks: post: deckid not found")
		return
	}

	accountscards := SE.AccountsCards.Cache[account.Username]
	if accountscards == nil {
		log.Error("mydecks: post: accountscards missing")
		return
	}

	for cardid, count := range deck.Cards {
		if accountscards[cardid] == nil {
			log.Add("DeckId", deck.Id).Error("mydecks: post: cardid not in collection")
			return
		} else if maxCount := len(accountscards[cardid]); maxCount < count {
			log.New().Add("Username", account.Username).Add("CardId", cardid).Add("RequestCount", count).Add("MaxCount", maxCount)
			log.Warn("mydecks: post: request more cards than in collection")
			accountsdecks[deck.Id].Cards[cardid] = maxCount
		} else if count == 0 {
			delete(accountsdecks[deck.Id].Cards, cardid)
		} else {
			accountsdecks[deck.Id].Cards[cardid] = count
		}
	}

	accountsdecks[deck.Id].Name = deck.Name
	accountsdecks[deck.Id].Register = time.Now()

	if err := SE.AccountsDecks.Delete(account.Username, deck.Id); err != nil {
		log.Add("Error", err).Error("mydecks: post: delete old deck")
		return
	} else if err := SE.AccountsDecks.Insert(account.Username, deck.Id); err != nil {
		log.Add("Error", err).Error("mydecks: post: insert new deck")
		return
	}

	MakeDeckJson(accountsdecks[deck.Id]).Write(w)
	log.Add("DeckId", deck.Id).Add("Name", deck.Name).Add("Cards", accountsdecks[deck.Id].Cards).Debug("mydecks: post success")
}
