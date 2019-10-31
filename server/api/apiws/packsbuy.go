package apiws

import (
	"github.com/zachtaylor/7elements/gencardpack"
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
)

func PacksBuy(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Root.Logger.New().Add("User", socket.GetUser()).Tag("apiws/packsbuy")
		session := socket.Session
		if session == nil {
			log.Warn("not logged in")
			pushError(socket, "not logged in")
			return
		}

		account, err := rt.Root.Accounts.Find(session.Name())
		if account == nil {
			log.Add("Error", err).Error("account missing")
			pushError(socket, "account missing")
			return
		}

		acs, err := rt.Root.AccountsCards.Get(account.Username)
		if acs == nil {
			log.Add("Error", err).Error("account collection missing")
			pushError(socket, "account collection missing")
			return
		}

		packid := m.Data.GetI("packid")
		if packid < 1 {
			log.Error("packid missing")
			pushError(socket, "packid missing")
			return
		}
		log.Add("PackID", packid)

		pack, err := rt.Root.Packs.Get(packid)
		if pack == nil {
			log.Add("Error", err).Error("pack missing")
			pushError(socket, "pack missing")
			return
		}

		if account.Coins < pack.Cost {
			log.Warn("insufficient")
			pushError(socket, "you don't have enough coins")
			return
		}

		account.Coins -= pack.Cost
		rt.Root.Accounts.UpdateCoins(account)

		cards := gencardpack.NewPack(rt.Root, account.Username, pack)
		for _, card := range cards {
			if err := rt.Root.AccountsCards.InsertCard(card); err != nil {
				log.Add("Error", err).Error("insertcard")
				pushError(socket, "500 internal server error")
				return
			}
		}
		log.Add("Cards", cards).Info()
		pushJSON(socket, "/data/myaccount", rt.Root.AccountJSON(account.Username))
	})
}
