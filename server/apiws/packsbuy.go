package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/runtime"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func PacksBuy(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Logger.Add("Socket", socket.ID())
		if len(socket.Name()) < 1 {
			log.Warn("anon pack buy")
			socket.WriteSync(wsout.ErrorJSON("vii: purchase", "you must log in to chat"))
			return
		}
		log = log.Add("Name", socket.Name())

		account := rt.Accounts.Get(socket.Name())
		if account == nil {
			log.Warn("account missing")
			socket.WriteSync(wsout.ErrorJSON("vii: purchase", "you must log in to chat"))
			return
		}
		log = log.Add("Session", account.SessionID)

		var packid int
		if packidbuff, _ := m.Data["packid"].(float64); packidbuff < 1 {
			log.Add("packid", m.Data["packid"]).Add("type", reflect.TypeOf(m.Data["packid"])).Warn("packid missing")
			socket.WriteSync(wsout.ErrorJSON("vii: purchase", "packid missing"))
			return
		} else {
			packid = int(packidbuff)
		}
		log = log.Add("Pack", packid)

		pack := rt.Packs[packid]
		if pack == nil {
			log.Warn("pack missing")
			socket.WriteSync(wsout.ErrorJSON("vii: purchase", "internal error"))
			return
		}

		if account.Coins < pack.Cost {
			log.Add("Coins", account.Coins).Add("Cost", pack.Cost).Warn("insufficient")
			socket.WriteSync(wsout.ErrorJSON("vii: purchase", "not enough coins"))
			return
		}

		account.Coins -= pack.Cost
		accounts.UpdateCoins(rt.DB, account)

		cardIDs := pack.NewPack()
		log = log.Add("Cards", cardIDs)

		if err := accounts.InsertCards(rt.DB, account, cardIDs); err != nil {
			log.Add("Error", err).Error("insertcard")
			socket.WriteSync(wsout.ErrorJSON("vii: purchase", "internal error"))
			return
		}

		for _, cardID := range cardIDs {
			account.Cards[cardID]++
		}

		log.Info()
		socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
	})
}
