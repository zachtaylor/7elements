package apiws

import (
	"reflect"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/websocket"
)

func PacksBuy(server internal.Server) websocket.MessageHandler {
	return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := server.Log().Add("Socket", socket.ID())

		user := server.Users().GetWebsocket(socket)
		if user == nil {
			log.Warn("no session")
			socket.Write(websocket.MessageText, out.Error("vii", "you must log in to log out"))
			return
		}
		log = log.Add("Session", user.Session())

		account := server.Accounts().Get(user.Session().Name())
		if account == nil {
			log.Warn("account missing")
			socket.Write(websocket.MessageText, out.Error("vii: purchase", "you must log in to chat"))
			return
		}

		var packid int
		if packidbuff, _ := m.Data["packid"].(float64); packidbuff < 1 {
			log.Add("packid", m.Data["packid"]).Add("type", reflect.TypeOf(m.Data["packid"])).Warn("packid missing")
			socket.Write(websocket.MessageText, out.Error("vii: purchase", "packid missing"))
			return
		} else {
			packid = int(packidbuff)
		}
		log = log.Add("Pack", packid)

		pack := server.Content().Packs()[packid]
		if pack == nil {
			log.Warn("pack missing")
			socket.Write(websocket.MessageText, out.Error("vii: purchase", "internal error"))
			return
		}

		if account.Coins < pack.Cost {
			log.Add("Coins", account.Coins).Add("Cost", pack.Cost).Warn("insufficient")
			socket.Write(websocket.MessageText, out.Error("vii: purchase", "not enough coins"))
			return
		}

		account.Coins -= pack.Cost
		accounts.UpdateCoins(server.DB(), account)

		cardIDs := pack.NewPack()
		log = log.Add("Cards", cardIDs)

		if err := accounts.InsertCards(server.DB(), account, cardIDs); err != nil {
			log.Add("Error", err).Error("insertcard")
			socket.Write(websocket.MessageText, out.Error("vii: purchase", "internal error"))
			return
		}

		for _, cardID := range cardIDs {
			account.Cards[cardID]++
		}

		log.Info()
		socket.WriteMessage(websocket.NewMessage("/myaccount", account.Data()))
	})
}
