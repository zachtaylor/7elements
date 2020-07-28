package apiws

import (
	"github.com/zachtaylor/7elements/out"
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/http/websocket"
)

func PacksBuy(rt *runtime.T) websocket.Handler {
	const msgsrc = "purchase"
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if socket.Session == nil {
			out.Error(socket, msgsrc, "session required")
			return
		}
		log := rt.Log().Add("User", socket.Session.Name()).Tag("apiws/packsbuy")
		a, err := rt.Accounts.Get(socket.Session.Name())
		if a == nil {
			log.Add("Error", err).Error("account missing")
			out.Error(socket, msgsrc, "account missing")
			return
		}

		packid := m.Data.GetI("packid")
		if packid < 1 {
			log.Error("packid missing")
			out.Error(socket, msgsrc, "packid missing")
			return
		}
		log.Add("PackID", packid)

		pack, err := rt.Packs.Get(packid)
		if pack == nil {
			log.Add("Error", err).Error("pack missing")
			out.Error(socket, msgsrc, "pack missing")
			return
		}

		if a.Coins < pack.Cost {
			log.Warn("insufficient")
			out.Error(socket, msgsrc, "you don't have enough coins")
			return
		}

		a.Coins -= pack.Cost
		rt.Accounts.UpdateCoins(a)

		cardIDs := pack.NewPack()

		if err := rt.Accounts.InsertCards(a, cardIDs); err != nil {
			log.Add("Error", err).Error("insertcard")
			out.Error(socket, msgsrc, "500 internal server error")
			return
		}

		for _, cardID := range cardIDs {
			a.Cards[cardID]++
		}

		log.Add("Cards", cardIDs).Info()
		socket.Send("/myaccount", rt.AccountJSON(a))
	})
}
