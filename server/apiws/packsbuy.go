package apiws

import (
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func PacksBuy(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if socket.Session == nil {
			socket.Message("/error", cast.JSON{
				"error": "session required",
			})
			return
		}
		log := rt.Runtime.Root.Logger.New().Add("User", socket.Session.Name()).Tag("apiws/packsbuy")
		a, err := rt.Runtime.Root.Accounts.Find(socket.Session.Name())
		if a == nil {
			log.Add("Error", err).Error("account missing")
			socket.Message("/error", cast.JSON{
				"error": "account missing",
			})
			return
		}

		acs, err := rt.Runtime.Root.AccountsCards.Get(a.Username)
		if acs == nil {
			log.Add("Error", err).Error("account collection missing")
			socket.Message("/error", cast.JSON{
				"error": "account collection missing",
			})
			return
		}

		packid := m.Data.GetI("packid")
		if packid < 1 {
			log.Error("packid missing")
			socket.Message("/error", cast.JSON{
				"error": "packid missing",
			})
			return
		}
		log.Add("PackID", packid)

		pack, err := rt.Runtime.Root.Packs.Get(packid)
		if pack == nil {
			log.Add("Error", err).Error("pack missing")
			socket.Message("/error", cast.JSON{
				"error": "pack missing",
			})
			return
		}

		if a.Coins < pack.Cost {
			log.Warn("insufficient")
			socket.Message("/error", cast.JSON{
				"error": "you don't have enough coins",
			})
			return
		}

		a.Coins -= pack.Cost
		rt.Runtime.Root.Accounts.UpdateCoins(a)

		cardIDs := pack.NewPack()
		now := cast.Now()
		for _, cardID := range cardIDs {
			card := &account.Card{
				Username: a.Username,
				ProtoID:  cardID,
				Register: now,
			}
			if err := rt.Runtime.Root.AccountsCards.InsertCard(card); err != nil {
				log.Add("Error", err).Error("insertcard")
				socket.Message("/error", cast.JSON{
					"error": "500 internal server error",
				})
				return
			}
		}
		log.Add("Cards", cardIDs).Info()
		socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(a))
	})
}
