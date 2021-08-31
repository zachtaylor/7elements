package apihttp

// import (
// 	"net/http"
// 	"time"

// 	"github.com/zachtaylor/7elements"
// 	"github.com/zachtaylor/7elements/gencardpack"
// 	zhttp "ztaylor.me/http"
// 	"ztaylor.me/js"
// 	"ztaylor.me/log"
// )

// var OpenPackJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	log := log.Add("Addr", r.RemoteAddr)

// 	if r.Method != "GET" {
// 		w.WriteHeader(404)
// 		log.Add("Method", r.Method).Warn("openpack.json: only GET allowed")
// 		return
// 	}

// 	session, err := zhttp.ReadRequestCookie(r)
// 	if session == nil {
// 		if err != nil {
// 			log.Add("Error", err)
// 		}
// 		w.WriteHeader(400)
// 		w.Write([]byte("session missing"))
// 		log.Error("openpack.json: session missing")
// 		return
// 	}

// 	log.Add("Username", session.Username)
// 	account := vii.AccountService.Test(session.Username)
// 	if account == nil {
// 		w.WriteHeader(500)
// 		log.Error("openpack.json: account missing")
// 		return
// 	}

// 	if account.Packs < 1 {
// 		w.WriteHeader(500)
// 		w.Write([]byte("no packs"))
// 		log.Warn("openpack.json: no packs to open")
// 		return
// 	}

// 	account.Packs--
// 	startTime := time.Now()
// 	if err := vii.AccountService.UpdatePacks(account); err != nil {
// 		w.Write([]byte("error opening pack"))
// 		log.Add("Error", err).Error("openpack.json: error opening pack")
// 		return
// 	}

// 	js.Object{
// 		"cards": carddata,
// 	}.Write(w)
// 	log.Add("PacksRemaining", account.Packs).Add("Timer", time.Now().Sub(startTime)).Add("Cards", carddata).Info("openpack")
// })
