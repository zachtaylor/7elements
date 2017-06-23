package openpack

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
	"github.com/cznic/mathutil"
	"net/http"
	"time"
)

var cardRandomFC32s [7]*mathutil.FC32

func ShufflePacks() {
	for i := 0; i < len(cardRandomFC32s); i++ {
		cardRandomFC32s[i], _ = mathutil.NewFC32(1, len(SE.Cards.Cache), true)
		cardRandomFC32s[i].Seed(time.Now().Unix() + int64(17*i))
	}
}

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "GET" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("openpack: only GET allowed")
		return
	}

	session, err := sessionman.ReadRequestCookie(r)
	if session == nil {
		if err != nil {
			sessionman.EraseSessionId(w)
			log.Add("Error", err)
		}
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("openpack: session missing")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("openpack: account missing")
		return
	}

	accountcards, err := serverutil.GetAccountsCards(account.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("openpack: collection")
		return
	}

	accountpacks, err := serverutil.GetAccountsPacks(account.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("openpack: packs")
		return
	}

	if len(accountpacks) < 1 {
		w.WriteHeader(500)
		w.Write([]byte("no packs"))
		log.Add("Error", err).Error("openpack: no packs to open")
		return
	}

	pack := accountpacks[0]
	accountpacks = accountpacks[1:]
	SE.AccountsPacks.Cache[account.Username] = accountpacks
	register := time.Now()

	carddata := make([]uint, 0)
	j := json.Json{
		"register": pack.Register.String(),
	}

	for _, cardRandomFC32 := range cardRandomFC32s {
		cardid := uint(cardRandomFC32.Next())
		carddata = append(carddata, cardid)
		accountcard := &SE.AccountCard{
			Username: account.Username,
			Card:     cardid,
			Register: register,
		}

		if list := accountcards[cardid]; list != nil {
			accountcards[cardid] = append(list, accountcard)
		} else {
			accountcards[cardid] = []*SE.AccountCard{accountcard}
		}
	}

	j["cards"] = carddata
	j.Write(w)

	log.Add("Pack", pack).Add("PacksRemaining", len(accountpacks)).Add("Register", pack.Register.String()).Info("openpack: success")
})
