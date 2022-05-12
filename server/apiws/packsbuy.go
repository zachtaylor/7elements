package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func PacksBuy(server internal.Server) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		if len(socket.SessionID()) < 1 {
			log.Warn("no session")
			socket.Write(wsout.Error("vii", "no user"))
			return
		}
		log = log.Add("Session", socket.SessionID())

		user, _, err := server.GetUserManager().GetSession(socket.SessionID())
		if user == nil {
			log.Add("Error", err).Error("user missing")
			socket.Write(wsout.Error("vii", "internal error"))
			return
		}
		log = log.Add("Username", user.Name())

		account := server.GetAccounts().Get(user.Name())
		if account == nil {
			log.Warn("account missing")
			socket.WriteSync(wsout.Error("vii: purchase", "you must log in to chat"))
			return
		}

		var packid int
		if packidbuff, _ := m.Data["packid"].(float64); packidbuff < 1 {
			log.Add("packid", m.Data["packid"]).Add("type", reflect.TypeOf(m.Data["packid"])).Warn("packid missing")
			socket.WriteSync(wsout.Error("vii: purchase", "packid missing"))
			return
		} else {
			packid = int(packidbuff)
		}
		log = log.Add("Pack", packid)

		pack := server.GetGameVersion().GetPacks()[packid]
		if pack == nil {
			log.Warn("pack missing")
			socket.WriteSync(wsout.Error("vii: purchase", "internal error"))
			return
		}

		if account.Coins < pack.Cost {
			log.Add("Coins", account.Coins).Add("Cost", pack.Cost).Warn("insufficient")
			socket.WriteSync(wsout.Error("vii: purchase", "not enough coins"))
			return
		}

		account.Coins -= pack.Cost
		accounts.UpdateCoins(server.GetDB(), account)

		cardIDs := pack.NewPack()
		log = log.Add("Cards", cardIDs)

		if err := accounts.InsertCards(server.GetDB(), account, cardIDs); err != nil {
			log.Add("Error", err).Error("insertcard")
			socket.WriteSync(wsout.Error("vii: purchase", "internal error"))
			return
		}

		for _, cardID := range cardIDs {
			account.Cards[cardID]++
		}

		log.Info()
		socket.WriteSync(wsout.MyAccount(account.Data()).EncodeToJSON())
	})
}
